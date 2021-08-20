package xalert

import (
	"context"
	"fmt"
	"strings"

	"github.com/ashwanthkumar/slack-go-webhook"

	"github.com/tikivn/ultrago/env"
	"github.com/tikivn/ultrago/xlogaff"
)

type SlackAddress struct {
	Username  string
	Channel   string
	IconEmoji string
}

func (s SlackAddress) SendSlackMessage(ctx context.Context, message string) {
	ctx, logger := xlogaff.GetLogger(ctx)
	logger.Info(message)
	text := fmt.Sprintf("[%s]\n%s", strings.ToUpper(env.ENV), message)
	s.slackAlert(ctx, text)
}

func (s SlackAddress) SendSlackError(ctx context.Context, message string) {
	ctx, logger := xlogaff.GetLogger(ctx)
	logger.Error(message)
	text := fmt.Sprintf("[%s][ERROR]\n%s", strings.ToUpper(env.ENV), message)
	s.slackAlert(ctx, text)
}

func (s SlackAddress) SendSlackWarn(ctx context.Context, message string) {
	ctx, logger := xlogaff.GetLogger(ctx)
	logger.Warn(message)
	text := fmt.Sprintf("[%s][WARN]\n%s", strings.ToUpper(env.ENV), message)
	s.slackAlert(ctx, text)
}

func (s SlackAddress) slackAlert(ctx context.Context, message string) {
	ctx, logger := xlogaff.GetLogger(ctx)
	if env.SLACK_WEBHOOK_URL != "" {
		payload := slack.Payload{
			Text:        message,
			Username:    s.Username,
			Channel:     s.Channel,
			IconEmoji:   s.IconEmoji,
			Attachments: nil,
		}
		err := slack.Send(env.SLACK_WEBHOOK_URL, "", payload)
		if len(err) > 0 {
			logger.Errorf("slack send errors: %s", err)
			return
		}
	} else {
		logger.Errorf("missing slack webhook url")
	}
}
