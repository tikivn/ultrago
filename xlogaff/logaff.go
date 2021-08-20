package xlogaff

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var (
	logger = logrus.New()
)

func SetFormatter(formatter logrus.Formatter) {
	if logger != nil {
		logger.Formatter = formatter
	}
}

func SetLevel(level logrus.Level) {
	if logger != nil {
		logger.Level = level
	}
}

func NewDefaultFormatter() logrus.Formatter {
	formatter := new(logrus.TextFormatter)
	formatter.TimestampFormat = "2006-01-02 15:04:05"
	formatter.FullTimestamp = true
	return formatter
}

const (
	loggerKey  = "logger"
	sessionKey = "session"
	userKey    = "user"
	sessionID  = "session-id"
	userID     = "user-id"
	trackingID = "tracking-id"
)

type BaseUser interface {
	GetID() string
}

type BaseSession interface {
	GetID() string
	GetUserID() string
}

type Logger struct {
	*logrus.Entry
}

func GetNewLogger() *Logger {
	return &Logger{logrus.NewEntry(logger)}
}

func GetLogger(ctx context.Context) (context.Context, *Logger) {
	if log, exist := ctx.Value(loggerKey).(*Logger); exist && log.Entry != nil {
		return ctx, log
	} else {
		entryLogger := logger.WithField(trackingID, uuid.New().String())
		if session, exist := ctx.Value(sessionKey).(BaseSession); exist {
			if session.GetID() != "" {
				entryLogger = entryLogger.WithField(sessionID, session.GetID())
			}
			if session.GetUserID() != "" {
				entryLogger = entryLogger.WithField(userID, session.GetUserID())
			}
		}
		instance := &Logger{entryLogger}
		return context.WithValue(ctx, loggerKey, instance), instance
	}
}

func UpdateSessionToLogger(ctx context.Context) context.Context {
	if session, exist := ctx.Value(sessionKey).(BaseSession); exist {
		if log, exist := ctx.Value(loggerKey).(*Logger); exist && log.Entry != nil {
			entryLogger := log.Entry
			if session.GetID() != "" {
				entryLogger = entryLogger.WithField(sessionID, session.GetID())
			}
			if session.GetUserID() != "" {
				entryLogger = entryLogger.WithField(userID, session.GetUserID())
			}
			instance := &Logger{entryLogger}
			return context.WithValue(ctx, loggerKey, instance)
		} else {
			instance := &Logger{logger.WithField(trackingID, uuid.New().String())}
			return context.WithValue(ctx, loggerKey, instance)
		}
	}
	return ctx
}

func UpdateUserToLogger(ctx context.Context) context.Context {
	if user, exist := ctx.Value(userKey).(BaseUser); exist {
		if log, exist := ctx.Value(loggerKey).(*Logger); exist && log.Entry != nil {
			entryLogger := log.Entry
			if user.GetID() != "" {
				entryLogger = entryLogger.WithField(userID, user.GetID())
			}
			instance := &Logger{entryLogger}
			return context.WithValue(ctx, loggerKey, instance)
		} else {
			instance := &Logger{logger.WithField(trackingID, uuid.New().String())}
			return context.WithValue(ctx, loggerKey, instance)
		}
	}
	return ctx
}

func (l *Logger) SetTrackingID(id string) *Logger {
	return &Logger{l.Entry.WithField(trackingID, id)}
}

func (l *Logger) WithField(key string, value interface{}) *Logger {
	if v, err := GetJSON(value); err == nil {
		return &Logger{l.Entry.WithField(key, v)}
	}
	return &Logger{l.Entry.WithField(key, value)}
}

func (l *Logger) WithFields(fields logrus.Fields) *Logger {
	var (
		v   string
		err error
	)
	for key, value := range fields {
		if v, err = GetJSON(value); err == nil {
			fields[key] = v
		}
	}
	return &Logger{l.Entry.WithFields(fields)}
}

func marshalArgs(args ...interface{}) []interface{} {
	var (
		v   string
		err error
	)
	length := len(args)
	for i := 0; i < length; i++ {
		if v, err = GetJSON(args[i]); err == nil {
			args[i] = v
		}
	}
	return args
}

func (l *Logger) Log(level logrus.Level, args ...interface{}) {
	l.Entry.Log(level, marshalArgs(args...)...)
}

func (l *Logger) Logln(level logrus.Level, args ...interface{}) {
	l.Entry.Logln(level, marshalArgs(args...)...)
}

func (l *Logger) Logf(level logrus.Level, format string, args ...interface{}) {
	l.Entry.Logf(level, format, marshalArgs(args...)...)
}

func (l *Logger) Info(args ...interface{}) {
	l.Entry.Info(marshalArgs(args...)...)
}

func (l *Logger) Infoln(args ...interface{}) {
	l.Entry.Infoln(marshalArgs(args...)...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.Entry.Infof(format, marshalArgs(args...)...)
}

func (l *Logger) Debug(args ...interface{}) {
	l.Entry.Debug(marshalArgs(args...)...)
}

func (l *Logger) Debugln(args ...interface{}) {
	l.Entry.Debugln(marshalArgs(args...)...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.Entry.Debugf(format, marshalArgs(args...)...)
}

func (l *Logger) Trace(args ...interface{}) {
	l.Entry.Trace(marshalArgs(args...)...)
}

func (l *Logger) Traceln(args ...interface{}) {
	l.Entry.Traceln(marshalArgs(args...)...)
}

func (l *Logger) Tracef(format string, args ...interface{}) {
	l.Entry.Tracef(format, marshalArgs(args...)...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.Entry.Warn(marshalArgs(args...)...)
}

func (l *Logger) Warnln(args ...interface{}) {
	l.Entry.Errorln(marshalArgs(args...)...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.Entry.Warnf(format, marshalArgs(args...)...)
}

func (l *Logger) Error(args ...interface{}) {
	l.Entry.Error(marshalArgs(args...)...)
}

func (l *Logger) Errorln(args ...interface{}) {
	l.Entry.Errorln(marshalArgs(args...)...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Entry.Errorf(format, marshalArgs(args...)...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.Entry.Fatal(marshalArgs(args...)...)
}

func (l *Logger) Fatalln(args ...interface{}) {
	l.Entry.Fatalln(marshalArgs(args...)...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.Entry.Fatalf(format, marshalArgs(args...)...)
}

func (l *Logger) Panic(args ...interface{}) {
	l.Entry.Panic(marshalArgs(args...)...)
}

func (l *Logger) Panicln(args ...interface{}) {
	l.Entry.Panicln(marshalArgs(args...)...)
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.Entry.Panicf(format, marshalArgs(args...)...)
}

func GetJSON(value interface{}) (string, error) {
	switch value.(type) {
	case error, *error:
		return "", errors.New("json unsupported type: error")
	}
	k := reflect.ValueOf(value).Kind()
	switch k {
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Struct, reflect.Ptr:
		v, err := json.Marshal(value)
		return string(v), err
	}
	return "", errors.New(fmt.Sprintf("json unsupported type: %s", k.String()))
}
