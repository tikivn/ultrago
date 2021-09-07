package u_logger

import (
	"context"
	"errors"
	"fmt"
	"time"

	gorm "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

func NewGORMLogger(config gorm.Config) gorm.Interface {
	return &GORMLogger{
		Config:       config,
		InfoStr:      "%s\n[info] ",
		WarnStr:      "%s\n[warn] ",
		ErrStr:       "%s\n[error] ",
		TraceStr:     "%s\n[%.3fms] [rows:%v] %s",
		TraceWarnStr: "%s %s\n[%.3fms] [rows:%v] %s",
		TraceErrStr:  "%s %s\n[%.3fms] [rows:%v] %s",
	}
}

type GORMLogger struct {
	Config       gorm.Config
	InfoStr      string
	WarnStr      string
	ErrStr       string
	TraceStr     string
	TraceWarnStr string
	TraceErrStr  string
}

func (g *GORMLogger) LogMode(level gorm.LogLevel) gorm.Interface {
	return NewGORMLogger(gorm.Config{
		SlowThreshold:             time.Second,
		IgnoreRecordNotFoundError: true,
		LogLevel:                  level,
	})
}

func (g *GORMLogger) Printf(ctx context.Context, msg string, data ...interface{}) {
	ctx, logger := GetLogger(ctx)
	logger.Printf(msg, data...)
}

func (g *GORMLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if g.Config.LogLevel >= gorm.Info {
		g.Printf(ctx, g.InfoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

func (g *GORMLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if g.Config.LogLevel >= gorm.Warn {
		g.Printf(ctx, g.WarnStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

func (g *GORMLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if g.Config.LogLevel >= gorm.Error {
		g.Printf(ctx, g.ErrStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

func (g *GORMLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if g.Config.LogLevel > gorm.Silent {
		elapsed := time.Since(begin)
		switch {
		case err != nil && g.Config.LogLevel >= gorm.Error && (!errors.Is(err, gorm.ErrRecordNotFound) || !g.Config.IgnoreRecordNotFoundError):
			sql, rows := fc()
			if rows == -1 {
				g.Printf(ctx, g.TraceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				g.Printf(ctx, g.TraceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case elapsed > g.Config.SlowThreshold && g.Config.SlowThreshold != 0 && g.Config.LogLevel >= gorm.Warn:
			sql, rows := fc()
			slowLog := fmt.Sprintf("SLOW SQL >= %v", g.Config.SlowThreshold)
			if rows == -1 {
				g.Printf(ctx, g.TraceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				g.Printf(ctx, g.TraceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case g.Config.LogLevel == gorm.Info:
			sql, rows := fc()
			if rows == -1 {
				g.Printf(ctx, g.TraceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				g.Printf(ctx, g.TraceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		}
	}
}
