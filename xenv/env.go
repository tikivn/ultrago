package xenv

import "github.com/tikivn/ultrago/xenv_parser"

var (
	ENV                = xenv_parser.GetString("ENV", "dev")
	SLACK_WEBHOOK_URL  = xenv_parser.GetString("SLACK_WEBHOOK_URL", "")
	TELEGRAM_BOT_TOKEN = xenv_parser.GetString("TELEGRAM_BOT_TOKEN", "")
	TELEGRAM_CHANNELS  = xenv_parser.GetString("TELEGRAM_CHANNELS", "")
)
