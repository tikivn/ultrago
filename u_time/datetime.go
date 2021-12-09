package u_time

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

const (
	DateTimeFormat string = "2006-01-02T15:04:05"
	DateFormat     string = "2006-01-02"
	MonthFormat    string = "2006-01"
)

func HoChiMinhTz() (*time.Location, error) {
	loc, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	return loc, err
}

func MustHoChiMinhTz() *time.Location {
	loc, err := HoChiMinhTz()
	if err != nil {
		panic(err)
	}
	return loc
}

func ToDateTimeStr(t time.Time, loc *time.Location) string {
	return t.In(loc).Format(DateTimeFormat)
}

func ToDateStr(t time.Time, loc *time.Location) string {
	return t.In(loc).Format(DateFormat)
}

func ToMonthStr(t time.Time, loc *time.Location) string {
	return t.In(loc).Format(MonthFormat)
}

func Millis2Str(format string, millis int64, loc *time.Location) string {
	t := Millis2Time(millis)
	return t.In(loc).Format(format)
}

func Millis2Time(millis int64) time.Time {
	return time.Unix(0, int64(time.Duration(millis)*time.Millisecond))
}

func Unix2Time(unix int64) time.Time {
	return time.Unix(unix, 0)
}

func Time2Millis(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func Time2StartMonth(t time.Time) time.Time {
	res := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	return res
}

func Time2EndMonth(t time.Time) time.Time {
	res := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	res = res.AddDate(0, 1, -1)
	return res
}

func Str2Time(format string, text string, loc *time.Location) (*time.Time, error) {
	t, err := time.ParseInLocation(format, text, loc)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func Str2Millis(format string, text string, loc *time.Location) (int64, error) {
	t, err := time.ParseInLocation(format, text, loc)
	if err != nil {
		return 0, err
	}
	return Time2Millis(t), nil
}

func CurrentMillis() int64 {
	return Time2Millis(time.Now())
}

func MillisToStartOfDay(millis int64, loc *time.Location) int64 {
	t := Millis2Time(millis)
	tt := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc)
	return Time2Millis(tt)
}

func FromDateTime(dt time.Time) *Date {
	d := Date(dt)
	return &d
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

