package u_logger

import (
	"context"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var (
	// std is the name of the standard logger in stdlib `log`
	std = logrus.New()
)

const (
	LoggerTrackingKey = "logger-tracking"

	loggerKey  = "logger"
	sessionID  = "session-id"
	userID     = "user-id"
	trackingID = "tracking-id"
)

func WithFormatter(level logrus.Level) {
	formatter := new(logrus.TextFormatter)
	formatter.TimestampFormat = "2006-01-02 15:04:05"
	formatter.FullTimestamp = true
	std.Formatter = formatter
	if std != nil {
		std.Level = level
	}
}

func WithTracking(ctx context.Context, trackingInfo TrackingInfo) context.Context {
	if trackingInfo == nil {
		return ctx
	}

	if log, exist := ctx.Value(loggerKey).(*Logger); exist && log.Entry != nil {
		entryLogger := log.Entry
		if trackingInfo.GetSessionID() != "" {
			entryLogger = entryLogger.WithField(sessionID, trackingInfo.GetSessionID())
		}
		if trackingInfo.GetUserID() != "" {
			entryLogger = entryLogger.WithField(userID, trackingInfo.GetUserID())
		}
		instance := &Logger{entryLogger}
		return context.WithValue(ctx, loggerKey, instance)
	} else {
		instance := &Logger{std.WithField(trackingID, uuid.New().String())}
		return context.WithValue(ctx, loggerKey, instance)
	}
}

type TrackingInfo interface {
	GetSessionID() string
	GetUserID() string
}
