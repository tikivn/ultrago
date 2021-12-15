package u_time

import (
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

