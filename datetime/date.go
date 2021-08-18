package datetime

import (
	"fmt"
	"strings"
	"time"
)

type Date time.Time

const (
	timeLayout      = "2006-01-02"
	locationDefault = "Asia/Ho_Chi_Minh"
)

func (d *Date) UnmarshalJSON(b []byte) error {
	var (
		loc *time.Location
		err error
	)
	loc, err = time.LoadLocation(locationDefault)
	if err != nil {
		loc = time.Local
	}
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
	var (
		loc *time.Location
		err error
	)
	loc, err = time.LoadLocation(locationDefault)
	if err != nil {
		loc = time.Local
	}
	if d == nil {
		return time.Time{}
	}
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

func Today() *Date {
	var (
		loc *time.Location
		err error
	)
	loc, err = time.LoadLocation(locationDefault)
	if err != nil {
		loc = time.Local
	}
	now := time.Now()
	today := Date(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc))
	return &today
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
