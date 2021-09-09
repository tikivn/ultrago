package u_time

import (
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestToDateUtil(t *testing.T) {
	convey.Convey("TestToDateUtils", t, func() {
		date := Today()
		fmt.Println(date.ToString())
	})
}
