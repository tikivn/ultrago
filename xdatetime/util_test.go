package xdatetime

import (
	"fmt"
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
)

func TestToDateKey(t *testing.T) {
	var millis int64 = 1594784309000

	convey.Convey("TestToDateKey", t, func() {
		dateKey1, err := ToDateKey(millis, "UTC")
		convey.So(err, convey.ShouldBeNil)
		convey.So(dateKey1, convey.ShouldEqual, "2020-07-15")

		dateKey2, err := ToDateKey(millis, "Asia/Ho_Chi_Minh")
		convey.So(err, convey.ShouldBeNil)
		convey.So(dateKey2, convey.ShouldEqual, "2020-07-15")

		dateKey3, err := ToDateKey(millis, "Asia/Saigon")
		convey.So(err, convey.ShouldBeNil)
		convey.So(dateKey3, convey.ShouldEqual, "2020-07-15")

		parseDate, err := ParseDate("2020-08-01")
		convey.So(err, convey.ShouldBeNil)
		convey.So(parseDate.Year(), convey.ShouldEqual, 2020)
		convey.So(parseDate.Month(), convey.ShouldEqual, 8)
		convey.So(parseDate.Day(), convey.ShouldEqual, 1)

		parseDateTime, err := ParseDateTime("2020-08-01T10:50:05")
		convey.So(err, convey.ShouldBeNil)
		convey.So(parseDateTime.Year(), convey.ShouldEqual, 2020)
		convey.So(parseDateTime.Month(), convey.ShouldEqual, 8)
		convey.So(parseDateTime.Day(), convey.ShouldEqual, 1)
		convey.So(parseDateTime.Hour(), convey.ShouldEqual, 10)
		convey.So(parseDateTime.Minute(), convey.ShouldEqual, 50)
		convey.So(parseDateTime.Second(), convey.ShouldEqual, 05)

		parseDateTimeFromUnixTsString, err := ParseDateTimeToMillis("1598506427", false)
		convey.So(err, convey.ShouldBeNil)
		convey.So(parseDateTimeFromUnixTsString.Year(), convey.ShouldEqual, 2020)
		convey.So(parseDateTimeFromUnixTsString.Month(), convey.ShouldEqual, 8)
		convey.So(parseDateTimeFromUnixTsString.Day(), convey.ShouldEqual, 27)

		parseDateTimeFromUnixTs, err := ParseDateTimeToMillis(1598506427000, true)
		convey.So(err, convey.ShouldBeNil)
		convey.So(parseDateTimeFromUnixTs.Year(), convey.ShouldEqual, 2020)
		convey.So(parseDateTimeFromUnixTs.Month(), convey.ShouldEqual, 8)
		convey.So(parseDateTimeFromUnixTs.Day(), convey.ShouldEqual, 27)

		roundUpDateFromUnixTs, err := RoundUpDateFromUnixTs(1598506427021, "Asia/Ho_Chi_Minh")
		convey.So(err, convey.ShouldBeNil)
		convey.So(roundUpDateFromUnixTs, convey.ShouldEqual, 1598461200000)

		fmt.Println(time.Now().AddDate(0, 0, 3).Unix() * 1000)

		millis, err := ParseDateKeyToMillis("2020-11-18")
		convey.So(err, convey.ShouldBeNil)
		fmt.Println(millis)
	})
}
