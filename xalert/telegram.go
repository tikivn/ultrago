package xalert

import (
	"context"
	"strconv"
	"strings"
	"sync"
	"tikivn/ultrago/env"
	"tikivn/ultrago/logaff"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var telegramToken string
var channels []string

func init() {
	telegramToken = env.TELEGRAM_BOT_TOKEN
	channelStr := env.TELEGRAM_CHANNELS
	if channelStr != "" {
		channels = strings.Split(channelStr, ",")
	}
}

func SendTeleMessage(ctx context.Context, message string) {
	ctx, logger := logaff.GetLogger(ctx)

	if telegramToken == "" {
		logger.Errorf("empty telegram token")
		return
	} else if len(channels) == 0 {
		logger.Errorf("empty list channels")
		return
	}

	bot, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		logger.Errorf("fail to connect to telegram: %v", err)
		return
	}

	var wg sync.WaitGroup
	for _, channel := range channels {
		wg.Add(1)
		go func(channel string, wg *sync.WaitGroup) {
			defer wg.Done()

			channelId, err := strconv.ParseInt(channel, 10, 64)
			if err != nil {
				logger.Errorf("invalid telegram channel id: %v", err)
				return
			}
			msg := tgbotapi.NewMessage(channelId, message)
			msg.ParseMode = "markdown"
			if _, err = bot.Send(msg); err != nil {
				logger.Errorf("fail to send message to telegram channel %s: %v", channel, err)
			}
		}(channel, &wg)
	}

	wg.Wait()
}
