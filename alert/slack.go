package alert

import (
	"context"
	"fmt"
	"strings"

	"github.com/ashwanthkumar/slack-go-webhook"

	"github.com/tikivn/tially/internal/pkg/setting"
	"github.com/tikivn/tially/pkg/util/logaff"
)

var (
	slackWebhookUrl = setting.SLACK_WEBHOOK_URL
)

func SendSlackAlert(ctx context.Context, message string) {
	ctx, logger := logaff.GetLogger(ctx)
	logger.Error(message)
	text := fmt.Sprintf("[%s][ERROR]\n%s", strings.ToUpper(setting.ENV), message)
	slackAlert(ctx, text)
}

func SlackWarn(ctx context.Context, message string) {
	ctx, logger := logaff.GetLogger(ctx)
	logger.Warn(message)
	text := fmt.Sprintf("[%s][WARN]\n%s", strings.ToUpper(setting.ENV), message)
	slackAlert(ctx, text)
}

func slackAlert(ctx context.Context, message string) {
	ctx, logger := logaff.GetLogger(ctx)
	if slackWebhookUrl != "" {
		payload := slack.Payload{
			Text:        message,
			Username:    "Tially Guardian",
			Channel:     "#tially-alert",
			IconEmoji:   ":ami_khoc:",
			Attachments: nil,
		}
		err := slack.Send(slackWebhookUrl, "", payload)
		if len(err) > 0 {
			logger.Errorf("error: %s", err)
			return
		}
	} else {
		logger.Errorf("missing slack webhook url")
	}
}
