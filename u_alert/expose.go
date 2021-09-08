package u_alert

import (
	"sync"

	"github.com/tikivn/ultrago/u_env"
)

var (
	slackIns  *slack
	slackOnce sync.Once

	telegramIns  *telegram
	telegramOnce sync.Once
)

const (
	SLACK_WEBHOOK_URL string = "SLACK_WEBHOOK_URL"

	TELEGRAM_BOT_TOKEN string = "TELEGRAM_BOT_TOKEN"
	TELEGRAM_CHANNELS  string = "TELEGRAM_CHANNELS"
)

func Slack() *slack {
	if slackIns == nil {
		slackOnce.Do(func() {
			slackIns = &slack{
				webhookURL: u_env.GetString(SLACK_WEBHOOK_URL, ""),
			}
		})
	}
	return slackIns
}

func Telegram() *telegram {
	if telegramIns == nil {
		telegramOnce.Do(func() {
			telegramIns = &telegram{
				token:    u_env.GetString(TELEGRAM_BOT_TOKEN, ""),
				channels: u_env.GetArray(TELEGRAM_CHANNELS, ",", nil),
			}
		})
	}
	return telegramIns
}
