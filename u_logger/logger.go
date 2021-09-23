package u_logger

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func NewLogger() *Logger {
	entry := std.WithField(trackingID, uuid.New().String())
	return &Logger{entry}
}

// NewLoggerCtx override existed logger key in ctx
func NewLoggerCtx(ctx context.Context, trackingId string) (context.Context, *Logger) {
	if trackingId == "" {
		trackingId = uuid.New().String()
	}
	entryLogger := std.WithField(trackingID, trackingId)
	instance := &Logger{entryLogger}
	return context.WithValue(ctx, loggerKey, instance), instance
}

func GetLogger(ctx context.Context) (context.Context, *Logger) {
	if ctx == nil {
		ctx = context.Background()
	}

	log, ok := ctx.Value(loggerKey).(*Logger)
	if ok && log.Entry != nil {
		return ctx, log
	}

	// init logger
	log = NewLogger()
	return context.WithValue(ctx, loggerKey, log), log
}

type Logger struct {
	*logrus.Entry
}

func (l *Logger) GetTrackingID() string {
	id, ok := l.Data[trackingID]
	if !ok {
		return "no-tracking-id"
	} else {
		res, ok := id.(string)
		if !ok {
			return "invalid-tracking-id"
		}
		return res
	}
}

func (l *Logger) WithField(key string, value interface{}) *Logger {
	if v, err := getJson(value); err == nil {
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
		if v, err = getJson(value); err == nil {
			fields[key] = v
		}
	}
	return &Logger{l.Entry.WithFields(fields)}
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

func marshalArgs(args ...interface{}) []interface{} {
	var (
		v   string
		err error
	)
	length := len(args)
	for i := 0; i < length; i++ {
		if v, err = getJson(args[i]); err == nil {
			args[i] = v
		}
	}
	return args
}

func getJson(value interface{}) (string, error) {
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
