package url

import (
	"github.com/smartystreets/goconvey/convey"
	"net/url"
	"testing"
)

func TestUrlValidator(t *testing.T) {
	convey.Convey("TestUrlValidator", t, func() {
		var (
			u   *url.URL
			err error
		)

		u, err = ValidateURL("http:::/not.valid/a//a??a?b=&&c#hi")
		convey.So(u, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)

		u, err = ValidateURL("http//google.com")
		convey.So(u, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)

		u, err = ValidateURL("google.com")
		convey.So(u, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)

		u, err = ValidateURL("/foo/bar")
		convey.So(u, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)

		u, err = ValidateURL("http://")
		convey.So(u, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)

		u, err = ValidateURL("http://google.com")
		convey.So(u, convey.ShouldNotBeNil)
		convey.So(err, convey.ShouldBeNil)

		u, err = ValidateURL("http://192.158.0.1:8000")
		convey.So(u, convey.ShouldNotBeNil)
		convey.So(err, convey.ShouldBeNil)

		u, err = ValidateURL("http://192.168.0.1")
		convey.So(u, convey.ShouldNotBeNil)
		convey.So(err, convey.ShouldBeNil)

		u, err = ValidateURL("https://tiki.vn/apple-macbook-air-2020-m1-13-inchs-apple-m1-16gb-256gb-hang-chinh-hang-p88231354.html?src=bestseller-page")
		convey.So(u, convey.ShouldNotBeNil)
		convey.So(err, convey.ShouldBeNil)

		u, err = ValidateURL("https%3A%2F%2Ftiki.vn%2Fapple-macbook-air-2020-m1-13-inchs-apple-m1-16gb-256gb-hang-chinh-hang-p88231354.html%3Fsrc%3Dbestseller-page")
		convey.So(u, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)
	})
}
