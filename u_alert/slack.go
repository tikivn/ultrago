package u_alert

import (
	"context"
	"fmt"
	"github.com/tikivn/ultrago/u_env"
	"strings"

	"github.com/ashwanthkumar/slack-go-webhook"
	env "github.com/tikivn/ultrago/u_env"
	logaff "github.com/tikivn/ultrago/u_logaff"
)

type SlackAddress struct {
	Username  string
	Channel   string
	IconEmoji string
}

func (s SlackAddress) SendSlackMessage(ctx context.Context, message string) {
	ctx, logger := logaff.GetLogger(ctx)
	logger.Info(message)
	text := fmt.Sprintf("[%s]\n%s", strings.ToUpper(u_env.ENV), message)
	s.slackAlert(ctx, text)
}

func (s SlackAddress) SendSlackError(ctx context.Context, message string) {
	ctx, logger := logaff.GetLogger(ctx)
	logger.Error(message)
	text := fmt.Sprintf("[%s][ERROR]\n%s", strings.ToUpper(u_env.ENV), message)
	s.slackAlert(ctx, text)
}

func (s SlackAddress) SendSlackWarn(ctx context.Context, message string) {
	ctx, logger := logaff.GetLogger(ctx)
	logger.Warn(message)
	text := fmt.Sprintf("[%s][WARN]\n%s", strings.ToUpper(u_env.ENV), message)
	s.slackAlert(ctx, text)
}

func (s SlackAddress) slackAlert(ctx context.Context, message string) {
	ctx, logger := logaff.GetLogger(ctx)
	if env.SLACK_WEBHOOK_URL != "" {
		payload := slack.Payload{
			Text:        message,
			Username:    s.Username,
			Channel:     s.Channel,
			IconEmoji:   s.IconEmoji,
			Attachments: nil,
		}
		err := slack.Send(u_env.SLACK_WEBHOOK_URL, "", payload)
		if len(err) > 0 {
			logger.Errorf("slack send errors: %s", err)
			return
		}
	} else {
		logger.Errorf("missing slack webhook url")
	}
}
