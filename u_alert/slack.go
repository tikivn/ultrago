package u_alert

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"

	slack_webhook "github.com/ashwanthkumar/slack-go-webhook"
	"github.com/tikivn/ultrago/u_env_parser"
	"github.com/tikivn/ultrago/u_logger"
)

var (
	slackIns  *slack
	slackOnce sync.Once
)

func Slack() *slack {
	if slackIns == nil {
		slackOnce.Do(func() {
			slackIns = &slack{
				webhookURL: u_env_parser.GetString("SLACK_WEBHOOK_URL", ""),
				botName:    u_env_parser.GetString("SLACK_BOT_NAME", "UltraGo"),
				channel:    u_env_parser.GetString("SLACK_CHANNEL", ""),
			}
		})
	}
	return slackIns
}

type slack struct {
	webhookURL string
	botName    string
	channel    string
}

func (s slack) env() string {
	return strings.ToUpper(u_env_parser.GetString("ENV", "DEV"))
}

func (s slack) Validate() error {
	if s.webhookURL == "" {
		return errors.New("empty slack webhook url")
	}
	if s.channel == "" {
		return errors.New("empty slack channel")
	}
	return nil
}

func (s slack) SendSlackMessage(ctx context.Context, message string) {
	ctx, logger := u_logger.GetLogger(ctx)
	logger.Info(message)
	text := fmt.Sprintf("[%s]\n%s", s.env(), message)
	s.slackAlert(ctx, text)
}

func (s slack) SendSlackError(ctx context.Context, message string) {
	ctx, logger := u_logger.GetLogger(ctx)
	logger.Error(message)
	text := fmt.Sprintf("[%s][ERROR]\n%s", s.env(), message)
	s.slackAlert(ctx, text)
}

func (s slack) SendSlackWarn(ctx context.Context, message string) {
	ctx, logger := u_logger.GetLogger(ctx)
	logger.Warn(message)
	text := fmt.Sprintf("[%s][WARN]\n%s", s.env(), message)
	s.slackAlert(ctx, text)
}

func (s slack) slackAlert(ctx context.Context, message string) {
	ctx, logger := u_logger.GetLogger(ctx)
	err := s.Validate()
	if err != nil {
		logger.Errorf(err.Error())
		return
	}

	payload := slack_webhook.Payload{
		Text:     message,
		Username: s.botName,
		Channel:  s.channel,
	}
	if slackErr := slack_webhook.Send(s.webhookURL, "", payload); len(slackErr) > 0 {
		logger.Errorf("slack send errors: %s", err)
		return
	}
}
