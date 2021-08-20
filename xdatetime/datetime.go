package xdatetime

import (
	"fmt"
	"strings"
	"time"
)

const (
	datetimeLayout = "2006-01-02 15:04:05"
)

type DateTime time.Time

func (d *DateTime) UnmarshalJSON(b []byte) error {
	loc := GetDefaultLocation()
	s := strings.Trim(string(b), "\"")
	t, err := time.ParseInLocation(datetimeLayout, s, loc)
	if err != nil {
		return err
	}
	*d = DateTime(t)
	return nil
}

func (d *DateTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", d.ToString())), nil
}

func (d *DateTime) ToTime() time.Time {
	if d == nil {
		return time.Time{}
	}
	loc := GetDefaultLocation()
	date, err := time.ParseInLocation(datetimeLayout, d.ToString(), loc)
	if err != nil {
		return time.Time{}
	}
	return date
}

func (d *DateTime) ToString() string {
	var t time.Time
	if d == nil {
		t = time.Time{}
	} else {
		t = time.Time(*d)
	}
	return fmt.Sprintf("%s", t.Format(datetimeLayout))
}

func (d *DateTime) ToEpochUnix() int64 {
	return d.ToTime().Unix() * 1000
}

func DateTimeNow() *DateTime {
	loc := GetDefaultLocation()
	now := time.Now()
	today := DateTime(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc))
	return &today
}

func GetDatetimeFromTime(t time.Time) *DateTime {
	dt := DateTime(t)
	return &dt
}
