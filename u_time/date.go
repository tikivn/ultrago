package u_time

import (
	"fmt"
	"strings"
	"time"
)

func FromDateTime(dt time.Time) *Date {
	d := Date(dt)
	return &d
}

func Today() *Date {
	loc, err := HoChiMinhTz()
	if err != nil {
		loc = time.Local
	}
	now := time.Now()
	today := Date(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc))
	return &today
}

type Date time.Time

func (d *Date) UnmarshalJSON(b []byte) error {

	loc, err := HoChiMinhTz()
	if err != nil {
		loc = time.Local
	}
	s := strings.Trim(string(b), "\"")
	t, err := time.ParseInLocation(DateFormat, s, loc)
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
	loc, err := HoChiMinhTz()
	if err != nil {
		loc = time.Local
	}
	if d == nil {
		return time.Time{}
	}
	date, err := time.ParseInLocation(DateFormat, d.ToString(), loc)
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
	return t.Format(DateFormat)
}

func (d *Date) ToEpochUnix() int64 {
	return d.ToTime().Unix() * 1000
}
