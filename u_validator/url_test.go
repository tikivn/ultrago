package u_validator

import (
	"net/url"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestVerifyURL(t *testing.T) {
	convey.Convey("TestVerifyURL", t, func() {
		var (
			u   *url.URL
			err error
		)

		u, err = VerifyURL("http:::/not.valid/a//a??a?b=&&c#hi")
		convey.So(u, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)

		u, err = VerifyURL("http//google.com")
		convey.So(u, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)

		u, err = VerifyURL("google.com")
		convey.So(u, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)

		u, err = VerifyURL("/foo/bar")
		convey.So(u, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)

		u, err = VerifyURL("http://")
		convey.So(u, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)

		u, err = VerifyURL("http://google.com")
		convey.So(u, convey.ShouldNotBeNil)
		convey.So(err, convey.ShouldBeNil)

		u, err = VerifyURL("http://192.158.0.1:8000")
		convey.So(u, convey.ShouldNotBeNil)
		convey.So(err, convey.ShouldBeNil)

		u, err = VerifyURL("http://192.168.0.1")
		convey.So(u, convey.ShouldNotBeNil)
		convey.So(err, convey.ShouldBeNil)

		u, err = VerifyURL("https://tiki.vn/apple-macbook-air-2020-m1-13-inchs-apple-m1-16gb-256gb-hang-chinh-hang-p88231354.html?src=bestseller-page")
		convey.So(u, convey.ShouldNotBeNil)
		convey.So(err, convey.ShouldBeNil)

		u, err = VerifyURL("https%3A%2F%2Ftiki.vn%2Fapple-macbook-air-2020-m1-13-inchs-apple-m1-16gb-256gb-hang-chinh-hang-p88231354.html%3Fsrc%3Dbestseller-page")
		convey.So(u, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)
	})
}
