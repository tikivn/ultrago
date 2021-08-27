package alarm

import (
	logaff "github.com/tikivn/ultrago/u_logaff"
	"time"
)

func New(hour int, minute int, internalDuration time.Duration) *alarm {
	return &alarm{
		atHour:           hour,
		atMinute:         minute,
		intervalDuration: internalDuration,
	}
}

type alarm struct {
	T *time.Timer

	atHour           int
	atMinute         int
	intervalDuration time.Duration
}

func (a *alarm) durationUntilNextAlarm(now time.Time) time.Duration {
	logger := logaff.GetNewLogger()
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	nextTick := time.Date(now.Year(), now.Month(), now.Day(), a.atHour, a.atMinute, 0, 0, loc)
	if nextTick.Before(now) {
		nextTick = nextTick.Add(a.intervalDuration)
	}

	duration := nextTick.Sub(now)
	logger.Infof("next alarm will be at %s", duration.String())
	return duration
}

func (a *alarm) Restart() {
	now := time.Now()
	duration := a.durationUntilNextAlarm(now)
	if a.T == nil {
		a.T = time.NewTimer(duration)
	} else {
		a.T.Reset(duration)
	}
}

func (a *alarm) Stop() {
	a.T.Stop()
}
