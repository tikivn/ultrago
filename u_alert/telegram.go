package u_alert

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/tikivn/ultrago/u_logger"
)

type telegram struct {
	token    string
	channels []string
}

func (t telegram) Validate() error {
	if t.token == "" {
		return errors.New("empty telegram token")
	}
	if len(t.channels) == 0 || t.channels[0] == "" {
		return errors.New("empty telegram channels")
	}
	return nil
}

func sendTele(ctx context.Context, wg *sync.WaitGroup, bot *tgbotapi.BotAPI, channel string, message string) {
	ctx, logger := u_logger.GetLogger(ctx)
	defer wg.Done()

	done := make(chan struct{})
	go func() {
		channelId, childErr := strconv.ParseInt(channel, 10, 64)
		if childErr != nil {
			logger.Errorf("invalid telegram channel id: %v", childErr)
			return
		}

		msg := tgbotapi.NewMessage(channelId, message)
		msg.ParseMode = "markdown"
		if _, childErr = bot.Send(msg); childErr != nil {
			logger.Errorf("fail to send message to telegram channel %s: %v", channel, childErr)
		}
		done <- struct{}{} // signal work is done
	}()

	select {
	case <-done:
		logger.Infof("send message to tele channel %s success", channel)
	case <-ctx.Done():
		logger.Infof("exit from ctx done")
	}
}

func (t telegram) SendMessage(ctx context.Context, message string) error {
	ctx, logger := u_logger.GetLogger(ctx)
	err := t.Validate()
	if err != nil {
		logger.Errorf(err.Error())
		return err
	}

	bot, err := tgbotapi.NewBotAPI(t.token)
	if err != nil {
		logger.Errorf("fail to connect to telegram: %v", err)
		return err
	}

	var wg sync.WaitGroup
	for idx := range t.channels {
		channel := t.channels[idx]
		wg.Add(1)
		go sendTele(ctx, &wg, bot, channel, message)
	}
	wg.Wait()
	return nil
}

func (t telegram) SendMessageMarkdown(ctx context.Context, message string) error {
	return t.SendMessage(ctx, t.escapeMarkdown(message))
}

/*
 * prepend '\' before special characters
 * https://core.telegram.org/bots/api#formatting-options
 */
func (t telegram) escapeMarkdown(message string) string {
	replaceChars := []string{"_", "*", "`", "["}
	for _, char := range replaceChars {
		message = strings.ReplaceAll(message, char, fmt.Sprintf("\\%s", char))
	}
	return message
}
