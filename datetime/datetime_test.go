package datetime

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

func TestToStartEndOfMonth(t *testing.T) {
	convey.Convey("TestToEndOfMonth", t, func() {
		january, _ := ParseDate("2021-01-03")
		february, _ := ParseDate("2021-02-03")
		march, _ := ParseDate("2021-03-03")
		april, _ := ParseDate("2021-04-03")
		may, _ := ParseDate("2021-05-03")
		june, _ := ParseDate("2021-06-03")
		july, _ := ParseDate("2021-07-03")
		august, _ := ParseDate("2021-08-03")
		september, _ := ParseDate("2021-09-03")
		october, _ := ParseDate("2021-10-03")
		november, _ := ParseDate("2021-11-03")
		december, _ := ParseDate("2021-12-03")

		startMonth1 := ToStartOfMonth(*january)
		startMonth2 := ToStartOfMonth(*february)
		startMonth3 := ToStartOfMonth(*march)
		startMonth4 := ToStartOfMonth(*april)
		startMonth5 := ToStartOfMonth(*may)
		startMonth6 := ToStartOfMonth(*june)
		startMonth7 := ToStartOfMonth(*july)
		startMonth8 := ToStartOfMonth(*august)
		startMonth9 := ToStartOfMonth(*september)
		startMonth10 := ToStartOfMonth(*october)
		startMonth11 := ToStartOfMonth(*november)
		startMonth12 := ToStartOfMonth(*december)
		convey.So(Time2Str(startMonth1), convey.ShouldEqual, "2021-01-01")
		convey.So(Time2Str(startMonth2), convey.ShouldEqual, "2021-02-01")
		convey.So(Time2Str(startMonth3), convey.ShouldEqual, "2021-03-01")
		convey.So(Time2Str(startMonth4), convey.ShouldEqual, "2021-04-01")
		convey.So(Time2Str(startMonth5), convey.ShouldEqual, "2021-05-01")
		convey.So(Time2Str(startMonth6), convey.ShouldEqual, "2021-06-01")
		convey.So(Time2Str(startMonth7), convey.ShouldEqual, "2021-07-01")
		convey.So(Time2Str(startMonth8), convey.ShouldEqual, "2021-08-01")
		convey.So(Time2Str(startMonth9), convey.ShouldEqual, "2021-09-01")
		convey.So(Time2Str(startMonth10), convey.ShouldEqual, "2021-10-01")
		convey.So(Time2Str(startMonth11), convey.ShouldEqual, "2021-11-01")
		convey.So(Time2Str(startMonth12), convey.ShouldEqual, "2021-12-01")

		endMonth1 := ToEndOfMonth(*january)
		endMonth2 := ToEndOfMonth(*february)
		endMonth3 := ToEndOfMonth(*march)
		endMonth4 := ToEndOfMonth(*april)
		endMonth5 := ToEndOfMonth(*may)
		endMonth6 := ToEndOfMonth(*june)
		endMonth7 := ToEndOfMonth(*july)
		endMonth8 := ToEndOfMonth(*august)
		endMonth9 := ToEndOfMonth(*september)
		endMonth10 := ToEndOfMonth(*october)
		endMonth11 := ToEndOfMonth(*november)
		endMonth12 := ToEndOfMonth(*december)
		convey.So(Time2Str(endMonth1), convey.ShouldEqual, "2021-01-31")
		convey.So(Time2Str(endMonth2), convey.ShouldEqual, "2021-02-28")
		convey.So(Time2Str(endMonth3), convey.ShouldEqual, "2021-03-31")
		convey.So(Time2Str(endMonth4), convey.ShouldEqual, "2021-04-30")
		convey.So(Time2Str(endMonth5), convey.ShouldEqual, "2021-05-31")
		convey.So(Time2Str(endMonth6), convey.ShouldEqual, "2021-06-30")
		convey.So(Time2Str(endMonth7), convey.ShouldEqual, "2021-07-31")
		convey.So(Time2Str(endMonth8), convey.ShouldEqual, "2021-08-31")
		convey.So(Time2Str(endMonth9), convey.ShouldEqual, "2021-09-30")
		convey.So(Time2Str(endMonth10), convey.ShouldEqual, "2021-10-31")
		convey.So(Time2Str(endMonth11), convey.ShouldEqual, "2021-11-30")
		convey.So(Time2Str(endMonth12), convey.ShouldEqual, "2021-12-31")
	})
}
