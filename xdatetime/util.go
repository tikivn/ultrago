package xdatetime

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

const (
	DateTimeFormat = "2006-01-02T15:04:05"
	DateFormat     = "2006-01-02"
	MonthFormat    = "2006-01"

	locationDefault = "Asia/Ho_Chi_Minh"
)

func ToStringFormat(millis int64, format string, timezone string) (string, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return "", err
	}
	date := time.Unix(0, millis*int64(time.Millisecond)).In(loc)
	return date.Format(format), nil
}

func ToDateKey(millis int64, timezone string) (string, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return "", err
	}
	date := time.Unix(0, millis*int64(time.Millisecond)).In(loc)
	return date.Format(DateFormat), nil
}

func ToMonthKey(millis int64, timezone string) (string, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return "", err
	}
	date := time.Unix(0, millis*int64(time.Millisecond)).In(loc)
	return date.Format(MonthFormat), nil
}

func LocalDateFromMillis(millis int64) string {
	return time.Unix(0, millis*int64(time.Millisecond)).Local().Format("2006-01-02")
}

func CurrentUnixMills() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func UnixMillsAfter(d time.Duration) int64 {
	return time.Now().Add(d).UnixNano() / int64(time.Millisecond)
}

func DurationBetween(fromMillis int64, toMillis int64) time.Duration {
	delta := int64(math.Abs(float64(toMillis - fromMillis)))
	return time.Duration(delta * int64(time.Millisecond))
}

func DurationFromNow(fromMillis int64) time.Duration {
	now := CurrentUnixMills()
	return DurationBetween(fromMillis, now)
}

func ToDayDate() string {
	return time.Now().Format(DateFormat)
}

func ParseDate(input string) (*time.Time, error) {
	var loc *time.Location
	loc, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		loc = time.Local
	}
	t, err := time.ParseInLocation(DateFormat, input, loc)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func ParseDateTime(input string) (*time.Time, error) {
	var loc *time.Location
	loc, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		loc = time.Local
	}
	t, err := time.ParseInLocation(DateTimeFormat, input, loc)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func ParseDateKeyToMillis(input string) (int64, error) {
	var loc *time.Location
	loc, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		loc = time.Local
	}
	t, err := time.ParseInLocation(DateFormat, input, loc)
	if err != nil {
		return 0, err
	}
	return t.Unix() * 1000, nil
}

func ParseDateTimeToMillis(input interface{}, isMillis bool) (*time.Time, error) {
	var t time.Time
	switch input.(type) {
	case string:
		i, err := strconv.ParseInt(input.(string), 10, 64)
		if err != nil {
			return nil, err
		}
		if isMillis {
			t = time.Unix(int64(i/1000), 0)
		} else {
			t = time.Unix(i, 0)
		}
	case int:
		if isMillis {
			t = time.Unix(int64(input.(int)/1000), 0)
		} else {
			t = time.Unix(int64(input.(int)), 0)
		}
	case int64:
		if isMillis {
			t = time.Unix(input.(int64)/1000, 0)
		} else {
			t = time.Unix(input.(int64), 0)
		}
	case float64:
		if isMillis {
			t = time.Unix(int64(input.(float64))/1000, 0)
		} else {
			t = time.Unix(int64(input.(float64)), 0)
		}
	}

	return &t, nil
}

func RoundUpDateFromUnixTs(input int64, timezone string) (int64, error) {
	i, err := strconv.ParseInt(fmt.Sprintf("%v", int64(input/1000)), 10, 64)
	if err != nil {
		return 0, err
	}
	t := time.Unix(i, 0)
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		panic(err)
	}
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc).Unix() * 1000, nil
}

func Millis2Time(millis int64) time.Time {
	return time.Unix(0, int64(time.Duration(millis)*time.Millisecond))
}

func Time2Millis(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
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

func FormatDurationMillis(millis int64, layout string) string {
	var duration = time.Duration(millis) * time.Millisecond
	return FormatDuration(duration, layout)
}

func GetDefaultLocation() *time.Location {
	var (
		loc *time.Location
		err error
	)
	loc, err = time.LoadLocation(locationDefault)
	if err != nil {
		loc = time.Local
	}
	return loc
}

func HoChiMinhTz() (*time.Location, error) {
	loc, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	return loc, err
}
