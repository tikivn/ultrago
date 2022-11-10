package u_alert

import (
	"strings"
	"sync"

	"github.com/tikivn/ultrago/u_env"
	"github.com/tikivn/ultrago/u_logger"
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

func init() {
	_, logger := u_logger.GetLogger(nil)
	slackEnabled := u_env.GetString(SLACK_WEBHOOK_URL, "") != ""
	telegramEnabled := u_env.GetString(TELEGRAM_BOT_TOKEN, "") != ""
	if slackEnabled {
		logger.Infof("Slack webhook is enabled")
	} else {
		logger.Warnf("Slack webhook is not enabled")
	}
	if telegramEnabled {
		logger.Infof("Telegram webhook is enabled")
	} else {
		logger.Warnf("Telegram webhook is not enabled")
	}
}

func InitCustomEnvVar(slackWebHook string, telegramBotToken string, telegramChannels string) {
	if slackWebHook != "" {
		slackOnce.Do(func() {
			slackIns = &slack{
				webhookURL: slackWebHook,
			}
		})
	}

	if telegramBotToken != "" && strings.TrimSpace(telegramChannels) != "" {
		telegramOnce.Do(func() {
			telegramIns = &telegram{
				token:    telegramBotToken,
				channels: strings.Split(telegramChannels, ","),
			}
		})
	}
}

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
