package u_time

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
)

func TestToDateKey(t *testing.T) {
	var millis int64 = 1594784309000

	convey.Convey("TestToDateKey", t, func() {
		os.Setenv("TZ", "UTC")
		dateKey1 := Millis2Str(DateFormat, millis, time.UTC)
		convey.So(dateKey1, convey.ShouldEqual, "2020-07-15")

		dateKey2 := Millis2Str(DateFormat, millis, MustHoChiMinhTz())
		convey.So(dateKey2, convey.ShouldEqual, "2020-07-15")

		parseDate, err := Str2Time(DateFormat, "2020-08-01", time.UTC)
		convey.So(err, convey.ShouldBeNil)
		convey.So(parseDate.Year(), convey.ShouldEqual, 2020)
		convey.So(parseDate.Month(), convey.ShouldEqual, 8)
		convey.So(parseDate.Day(), convey.ShouldEqual, 1)

		parseDateTime, err := Str2Time(DateTimeFormat, "2020-08-01T10:50:05", time.UTC)
		convey.So(err, convey.ShouldBeNil)
		convey.So(parseDateTime.Year(), convey.ShouldEqual, 2020)
		convey.So(parseDateTime.Month(), convey.ShouldEqual, 8)
		convey.So(parseDateTime.Day(), convey.ShouldEqual, 1)
		convey.So(parseDateTime.Hour(), convey.ShouldEqual, 10)
		convey.So(parseDateTime.Minute(), convey.ShouldEqual, 50)
		convey.So(parseDateTime.Second(), convey.ShouldEqual, 05)

		parseDateTimeFromUnixTsString := Unix2Time(1598506427)
		convey.So(parseDateTimeFromUnixTsString.Year(), convey.ShouldEqual, 2020)
		convey.So(parseDateTimeFromUnixTsString.Month(), convey.ShouldEqual, 8)
		convey.So(parseDateTimeFromUnixTsString.Day(), convey.ShouldEqual, 27)
		convey.So(parseDateTimeFromUnixTsString.Hour(), convey.ShouldEqual, 5)
		convey.So(parseDateTimeFromUnixTsString.Minute(), convey.ShouldEqual, 33)
		convey.So(parseDateTimeFromUnixTsString.Second(), convey.ShouldEqual, 47)

		parseDateTimeFromUnixTs := Millis2Time(1598506427000)
		convey.So(err, convey.ShouldBeNil)
		convey.So(parseDateTimeFromUnixTs.Year(), convey.ShouldEqual, 2020)
		convey.So(parseDateTimeFromUnixTs.Month(), convey.ShouldEqual, 8)
		convey.So(parseDateTimeFromUnixTs.Day(), convey.ShouldEqual, 27)
		convey.So(parseDateTimeFromUnixTsString.Hour(), convey.ShouldEqual, 5)
		convey.So(parseDateTimeFromUnixTsString.Minute(), convey.ShouldEqual, 33)
		convey.So(parseDateTimeFromUnixTsString.Second(), convey.ShouldEqual, 47)

		roundUpDateFromUnixTs := MillisToStartOfDay(1598506427021, MustHoChiMinhTz())
		convey.So(err, convey.ShouldBeNil)
		convey.So(roundUpDateFromUnixTs, convey.ShouldEqual, 1598461200000)

		fmt.Println(time.Now().AddDate(0, 0, 3).Unix() * 1000)

		millis, err := Str2Millis(DateFormat, "2020-11-18", time.UTC)
		convey.So(err, convey.ShouldBeNil)
		fmt.Println(millis)
	})
}

func TestToStartEndOfMonth(t *testing.T) {
	convey.Convey("TestToEndOfMonth", t, func() {
		january, _ := Str2Time(DateFormat, "2021-01-03", time.UTC)
		february, _ := Str2Time(DateFormat, "2021-02-03", time.UTC)
		march, _ := Str2Time(DateFormat, "2021-03-03", time.UTC)
		april, _ := Str2Time(DateFormat, "2021-04-03", time.UTC)
		may, _ := Str2Time(DateFormat, "2021-05-03", time.UTC)
		june, _ := Str2Time(DateFormat, "2021-06-03", time.UTC)
		july, _ := Str2Time(DateFormat, "2021-07-03", time.UTC)
		august, _ := Str2Time(DateFormat, "2021-08-03", time.UTC)
		september, _ := Str2Time(DateFormat, "2021-09-03", time.UTC)
		october, _ := Str2Time(DateFormat, "2021-10-03", time.UTC)
		november, _ := Str2Time(DateFormat, "2021-11-03", time.UTC)
		december, _ := Str2Time(DateFormat, "2021-12-03", time.UTC)

		startMonth1 := Time2StartMonth(*january)
		startMonth2 := Time2StartMonth(*february)
		startMonth3 := Time2StartMonth(*march)
		startMonth4 := Time2StartMonth(*april)
		startMonth5 := Time2StartMonth(*may)
		startMonth6 := Time2StartMonth(*june)
		startMonth7 := Time2StartMonth(*july)
		startMonth8 := Time2StartMonth(*august)
		startMonth9 := Time2StartMonth(*september)
		startMonth10 := Time2StartMonth(*october)
		startMonth11 := Time2StartMonth(*november)
		startMonth12 := Time2StartMonth(*december)
		convey.So(ToDateStr(startMonth1, time.UTC), convey.ShouldEqual, "2021-01-01")
		convey.So(ToDateStr(startMonth2, time.UTC), convey.ShouldEqual, "2021-02-01")
		convey.So(ToDateStr(startMonth3, time.UTC), convey.ShouldEqual, "2021-03-01")
		convey.So(ToDateStr(startMonth4, time.UTC), convey.ShouldEqual, "2021-04-01")
		convey.So(ToDateStr(startMonth5, time.UTC), convey.ShouldEqual, "2021-05-01")
		convey.So(ToDateStr(startMonth6, time.UTC), convey.ShouldEqual, "2021-06-01")
		convey.So(ToDateStr(startMonth7, time.UTC), convey.ShouldEqual, "2021-07-01")
		convey.So(ToDateStr(startMonth8, time.UTC), convey.ShouldEqual, "2021-08-01")
		convey.So(ToDateStr(startMonth9, time.UTC), convey.ShouldEqual, "2021-09-01")
		convey.So(ToDateStr(startMonth10, time.UTC), convey.ShouldEqual, "2021-10-01")
		convey.So(ToDateStr(startMonth11, time.UTC), convey.ShouldEqual, "2021-11-01")
		convey.So(ToDateStr(startMonth12, time.UTC), convey.ShouldEqual, "2021-12-01")

		endMonth1 := Time2EndMonth(*january)
		endMonth2 := Time2EndMonth(*february)
		endMonth3 := Time2EndMonth(*march)
		endMonth4 := Time2EndMonth(*april)
		endMonth5 := Time2EndMonth(*may)
		endMonth6 := Time2EndMonth(*june)
		endMonth7 := Time2EndMonth(*july)
		endMonth8 := Time2EndMonth(*august)
		endMonth9 := Time2EndMonth(*september)
		endMonth10 := Time2EndMonth(*october)
		endMonth11 := Time2EndMonth(*november)
		endMonth12 := Time2EndMonth(*december)
		convey.So(ToDateStr(endMonth1, time.UTC), convey.ShouldEqual, "2021-01-31")
		convey.So(ToDateStr(endMonth2, time.UTC), convey.ShouldEqual, "2021-02-28")
		convey.So(ToDateStr(endMonth3, time.UTC), convey.ShouldEqual, "2021-03-31")
		convey.So(ToDateStr(endMonth4, time.UTC), convey.ShouldEqual, "2021-04-30")
		convey.So(ToDateStr(endMonth5, time.UTC), convey.ShouldEqual, "2021-05-31")
		convey.So(ToDateStr(endMonth6, time.UTC), convey.ShouldEqual, "2021-06-30")
		convey.So(ToDateStr(endMonth7, time.UTC), convey.ShouldEqual, "2021-07-31")
		convey.So(ToDateStr(endMonth8, time.UTC), convey.ShouldEqual, "2021-08-31")
		convey.So(ToDateStr(endMonth9, time.UTC), convey.ShouldEqual, "2021-09-30")
		convey.So(ToDateStr(endMonth10, time.UTC), convey.ShouldEqual, "2021-10-31")
		convey.So(ToDateStr(endMonth11, time.UTC), convey.ShouldEqual, "2021-11-30")
		convey.So(ToDateStr(endMonth12, time.UTC), convey.ShouldEqual, "2021-12-31")
	})
}
