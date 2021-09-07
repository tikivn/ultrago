package u_logger

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockSession struct {
	ID     string
	UserID string
}

func (s *mockSession) GetSessionID() string {
	if s != nil {
		return s.ID
	}
	return ""
}

func (s *mockSession) GetUserID() string {
	if s != nil {
		return s.UserID
	}
	return ""
}

func TestGetLogger(t *testing.T) {

	t.Run("EmptyContext", func(t *testing.T) {
		ctx := context.Background()
		ctx, logger := GetLogger(ctx)
		if logger == nil || logger.Entry == nil || logger.Entry.Data == nil {
			t.Error("init logger error")
			return
		}
		if _, exist := logger.Entry.Data[trackingID]; !exist {
			t.Error("init logger not having trackingID")
		}
	})

	t.Run("ExistLoggerOnContext", func(t *testing.T) {
		ctx := context.Background()
		ctx, expected := GetLogger(ctx)
		ctx, logger := GetLogger(ctx)
		assert.Equal(t, expected, logger)
	})

	t.Run("ExistEmptyLoggerOnContext", func(t *testing.T) {
		emptyLogger := &Logger{}
		ctx := context.WithValue(context.Background(), loggerKey, emptyLogger)
		ctx, logger := GetLogger(ctx)
		assert.NotEqual(t, emptyLogger, logger)
	})

	t.Run("ExistSessionAndUserOnContext", func(t *testing.T) {
		session := "session-id"
		user := "user-id"

		ctx := context.WithValue(context.Background(), LoggerTrackingKey, &mockSession{
			ID:     session,
			UserID: user,
		})
		ctx, logger := GetLogger(ctx)
		if logger == nil || logger.Entry == nil || logger.Entry.Data == nil {
			t.Error("init logger error")
			return
		}
		if sessionID, exist := logger.Entry.Data[sessionID]; !exist {
			t.Error("init logger not having sessionID")
		} else if sessionID != session {
			t.Error("init logger with wrong sessionID")
		}
		if userID, exist := logger.Entry.Data[userID]; !exist {
			t.Error("init logger not having userID")
		} else if userID != user {
			t.Error("init logger with wrong userID")
		}
	})

	t.Run("NotExistUserOnContext", func(t *testing.T) {
		session := "session-id"

		ctx := context.WithValue(context.Background(), LoggerTrackingKey, &mockSession{
			ID: session,
		})
		ctx, logger := GetLogger(ctx)
		if logger == nil || logger.Entry == nil || logger.Entry.Data == nil {
			t.Error("init logger error")
			return
		}
		if _, exist := logger.Entry.Data[userID]; exist {
			t.Error("init logger having userID")
		}
	})

	t.Run("NotExistSessionOnContext", func(t *testing.T) {
		user := "user-id"

		ctx := context.WithValue(context.Background(), LoggerTrackingKey, &mockSession{
			UserID: user,
		})
		ctx, logger := GetLogger(ctx)
		if logger == nil || logger.Entry == nil || logger.Entry.Data == nil {
			t.Error("init logger error")
			return
		}
		if _, exist := logger.Entry.Data[sessionID]; exist {
			t.Error("init logger having sessionID")
		}
	})
}

func TestGetJSON(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "GetJSONArray",
			args: args{
				value: []int{1, 2, 3, 4, 5},
			},
			want:    "[1,2,3,4,5]",
			wantErr: false,
		},
		{
			name: "GetJSONMap",
			args: args{
				value: map[int]string{
					1: "element-1",
					2: "element-2",
				},
			},
			want:    "{\"1\":\"element-1\",\"2\":\"element-2\"}",
			wantErr: false,
		},
		{
			name: "GetJSONStruct",
			args: args{
				value: struct {
					Key   string `json:"key"`
					Value int    `json:"value"`
				}{
					Key:   "key",
					Value: 5,
				},
			},
			want:    "{\"key\":\"key\",\"value\":5}",
			wantErr: false,
		},
		{
			name: "GetJSONPointer",
			args: args{
				value: &struct {
					Key   string `json:"key"`
					Value int    `json:"value"`
				}{
					Key:   "key",
					Value: 5,
				},
			},
			want:    "{\"key\":\"key\",\"value\":5}",
			wantErr: false,
		},
		{
			name: "GetJSONString",
			args: args{
				value: "string",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "GetJSONInt",
			args: args{
				value: 5,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "GetJSONError",
			args: args{
				value: errors.New("new error"),
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getJson(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("getJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getJson() got = %v, want %v", got, tt.want)
			}
		})
	}
}
