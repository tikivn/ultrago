package u_time

import (
	"math"
	"time"
)

func DurationMillis(from int64, to int64) time.Duration {
	delta := int64(math.Abs(float64(to - from)))
	return time.Duration(delta * int64(time.Millisecond))
}

func DurationUtilNow(fromMillis int64) time.Duration {
	return DurationMillis(fromMillis, CurrentMillis())
}
