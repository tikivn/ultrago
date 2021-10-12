package u_alert

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tikivn/ultrago/u_env"
)

func TestSlack(t *testing.T) {
	t.Run("SendSuccess", func(t *testing.T) {
		ctx := context.Background()
		// need set env before run
		slackIns = &slack{
			webhookURL: u_env.GetString(SLACK_WEBHOOK_URL, ""),
		}
		assert.NotEmpty(t, slackIns.webhookURL)
		errs := Slack().slackAlert(ctx, fmt.Sprintf("slack test msg with formatter=%v", "test"))
		assert.Equal(t, 0, len(errs))
	})

	t.Run("SendFail_ValidateFail", func(t *testing.T) {
		ctx := context.Background()
		slackIns = &slack{
			webhookURL: "",
		}
		errs := Slack().slackAlert(ctx, fmt.Sprintf("slack test msg with formatter=%v", "test"))
		assert.Equal(t, 1, len(errs))
	})

	t.Run("SendFail_WebhookFail", func(t *testing.T) {
		ctx := context.Background()
		slackIns = &slack{
			webhookURL: "https://hooks.slack.com/services/abc",
		}
		errs := Slack().slackAlert(ctx, fmt.Sprintf("slack test msg with formatter=%v", "test"))
		assert.Equal(t, 1, len(errs))
	})

	t.Run("SendMessage_WithContext", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*80))
		defer cancel()

		slackIns = &slack{
			webhookURL: u_env.GetString(SLACK_WEBHOOK_URL, ""),
		}
		assert.NotEmpty(t, slackIns.webhookURL)
		errs := Slack().slackAlert(ctx, fmt.Sprintf("slack test msg with formatter=%v", "test"))
		assert.Equal(t, 1, len(errs))
	})
}
