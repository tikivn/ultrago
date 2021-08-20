package xdatetime

import (
	"fmt"
	"strings"
	"time"
)

const (
	timeLayout = "2006-01-02"
)

type Date time.Time

func (d *Date) UnmarshalJSON(b []byte) error {
	loc := GetDefaultLocation()
	s := strings.Trim(string(b), "\"")
	t, err := time.ParseInLocation(timeLayout, s, loc)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", d.ToString())), nil
}

func (d *Date) ToTime() time.Time {
	if d == nil {
		return time.Time{}
	}
	loc := GetDefaultLocation()
	date, err := time.ParseInLocation(timeLayout, d.ToString(), loc)
	if err != nil {
		return time.Time{}
	}
	return date
}

func (d *Date) ToString() string {
	var t time.Time
	if d == nil {
		t = time.Time{}
	} else {
		t = time.Time(*d)
	}
	return fmt.Sprintf("%s", t.Format(timeLayout))
}

func (d *Date) ToEpochUnix() int64 {
	return d.ToTime().Unix() * 1000
}

func DateNow() *Date {
	loc := GetDefaultLocation()
	now := time.Now()
	today := Date(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc))
	return &today
}

func GetDateFromTime(t time.Time) *Date {
	d := Date(t)
	return &d
}
