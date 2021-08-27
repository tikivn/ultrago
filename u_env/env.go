package u_env

import "github.com/tikivn/ultrago/xenv_parser"

var (
	ENV                = u_env_parser.GetString("ENV", "dev")
	SLACK_WEBHOOK_URL  = u_env_parser.GetString("SLACK_WEBHOOK_URL", "")
	TELEGRAM_BOT_TOKEN = u_env_parser.GetString("TELEGRAM_BOT_TOKEN", "")
	TELEGRAM_CHANNELS  = u_env_parser.GetString("TELEGRAM_CHANNELS", "")
)
