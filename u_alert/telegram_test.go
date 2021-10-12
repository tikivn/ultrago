package u_alert

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tikivn/ultrago/u_env"
)

func TestTelegram(t *testing.T) {
	t.Run("SendSuccess", func(t *testing.T) {
		ctx := context.Background()
		// need set env before run
		telegramIns = &telegram{
			token:    u_env.GetString(TELEGRAM_BOT_TOKEN, ""),
			channels: u_env.GetArray(TELEGRAM_CHANNELS, ",", nil),
		}
		assert.NotEmpty(t, telegramIns.token)
		assert.NotEmpty(t, telegramIns.channels)
		err := Telegram().SendMessage(ctx, fmt.Sprintf("telegram test msg with formatter=%v", "test"))
		assert.Nil(t, err)
	})

	t.Run("SendSuccess_Markdown", func(t *testing.T) {
		ctx := context.Background()
		message :=
			`Dear team,
Tiki AFF có program mới như sau:
Program A: 20.6% (2021-01-01 - 2021-01-10)
Ngân sách (VNĐ): 100000000000000000000
Get Link Program: 
https://test-url/get-link/program/testing-program-id
Thanks team.`
		telegramIns = &telegram{
			token:    u_env.GetString(TELEGRAM_BOT_TOKEN, ""),
			channels: u_env.GetArray(TELEGRAM_CHANNELS, ",", nil),
		}
		assert.NotEmpty(t, telegramIns.token)
		assert.NotEmpty(t, telegramIns.channels)
		err := Telegram().SendMessageMarkdown(ctx, message)
		assert.Nil(t, err)
	})

	t.Run("SendFail_ValidateFail", func(t *testing.T) {
		ctx := context.Background()
		telegramIns = &telegram{}
		err := Telegram().SendMessage(ctx, fmt.Sprintf("telegram test msg with formatter=%v", "test"))
		assert.NotNil(t, err)
	})

	t.Run("SendFail_WebhookFail", func(t *testing.T) {
		ctx := context.Background()
		telegramIns = &telegram{
			token:    "test-token",
			channels: []string{"test-channel"},
		}
		err := Telegram().SendMessage(ctx, fmt.Sprintf("telegram test msg with formatter=%v", "test"))
		assert.NotNil(t, err)
	})

	t.Run("SendMessage_WithContext", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*80))
		defer cancel()

		telegramIns = &telegram{
			token:    u_env.GetString(TELEGRAM_BOT_TOKEN, ""),
			channels: u_env.GetArray(TELEGRAM_CHANNELS, ",", nil),
		}
		assert.NotEmpty(t, telegramIns.token)
		assert.NotEmpty(t, telegramIns.channels)
		err := Telegram().SendMessage(ctx, fmt.Sprintf("telegram test msg with formatter=%v", "test"))
		assert.Nil(t, err)
	})
}
