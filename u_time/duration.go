package u_time

import (
	"fmt"
	"math"
	"time"
)

func DurationMillis(from int64, to int64) time.Duration {
	delta := int64(math.Abs(float64(to - from)))
	return time.Duration(delta * int64(time.Millisecond))
}

func DurationToNow(millis int64) time.Duration {
	return DurationMillis(millis, CurrentMillis())
}

func DurationFromNow(millis int64) time.Duration {
	return DurationMillis(CurrentMillis(), millis)
}

func FormatDuration(d time.Duration, layout string) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	return fmt.Sprintf(layout, h, m, s)
}