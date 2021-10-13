package u_alert

import (
	"context"
	"errors"
	"fmt"
	"strings"

	slack_webhook "github.com/slack-go/slack"
	"github.com/tikivn/ultrago/u_env"
	"github.com/tikivn/ultrago/u_logger"
)

type slack struct {
	webhookURL string
}

func (s slack) env() string {
	return strings.ToUpper(u_env.GetString("ENV", "DEV"))
}

func (s slack) Validate() error {
	if s.webhookURL == "" {
		return errors.New("empty slack webhook url")
	}
	return nil
}

func (s slack) SendMessage(ctx context.Context, message string) {
	ctx, logger := u_logger.GetLogger(ctx)
	logger.Info(message)
	text := fmt.Sprintf("[%s]\n%s", s.env(), message)
	s.slackAlert(ctx, text)
}

func (s slack) SendError(ctx context.Context, message string) {
	ctx, logger := u_logger.GetLogger(ctx)
	logger.Error(message)
	text := fmt.Sprintf("[%s][ERROR]\n%s", s.env(), message)
	s.slackAlert(ctx, text)
}

func (s slack) SendWarn(ctx context.Context, message string) {
	ctx, logger := u_logger.GetLogger(ctx)
	logger.Warn(message)
	text := fmt.Sprintf("[%s][WARN]\n%s", s.env(), message)
	s.slackAlert(ctx, text)
}

func (s slack) slackAlert(ctx context.Context, message string) []error {
	ctx, logger := u_logger.GetLogger(ctx)
	err := s.Validate()
	if err != nil {
		logger.Errorf(err.Error())
		return []error{err}
	}

	payload := &slack_webhook.WebhookMessage{
		Text: message,
	}
	if slackErr := slack_webhook.PostWebhookContext(ctx, s.webhookURL, payload); slackErr != nil {
		logger.Errorf("slack send errors: %v", slackErr)
		return []error{slackErr}
	}
	return nil
}
