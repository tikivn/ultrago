package u_alert

import (
	"context"
	"errors"
	"strconv"
	"sync"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/tikivn/ultrago/u_env_parser"
	"github.com/tikivn/ultrago/u_logger"
)

var (
	telegramIns  *telegram
	telegramOnce sync.Once
)

func Telegram() *telegram {
	if telegramIns == nil {
		telegramOnce.Do(func() {
			telegramIns = &telegram{
				token:    u_env_parser.GetString("TELEGRAM_BOT_TOKEN", ""),
				channels: u_env_parser.GetArray("TELEGRAM_CHANNELS", ",", nil),
			}
		})
	}
	return telegramIns
}

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

func (t telegram) SendTeleMessage(ctx context.Context, message string) {
	ctx, logger := u_logger.GetLogger(ctx)
	err := t.Validate()
	if err != nil {
		logger.Errorf(err.Error())
		return
	}

	bot, err := tgbotapi.NewBotAPI(t.token)
	if err != nil {
		logger.Errorf("fail to connect to telegram: %v", err)
		return
	}

	var wg sync.WaitGroup
	for idx := range t.channels {
		channel := t.channels[idx]
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
