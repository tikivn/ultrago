package datetime

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestToDateUtil(t *testing.T) {
	convey.Convey("TestToDateUtils", t, func() {
		date := Today()
		fmt.Println(date.ToString())
	})
}
