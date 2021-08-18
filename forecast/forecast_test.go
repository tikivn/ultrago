package forecast

import (
	"math"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestForecast(t *testing.T) {
	knownX := []float64{2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4}
	knownY := []float64{14139193, 14837335, 26850009, 16865211, 16871674, 14338330, 14373474, 15499791, 16940038, 17631929, 20140964, 16941284, 16398389, 13062832, 14358278, 34348199, 35140157, 37765482, 30296677, 17886126, 22967137, 32190795, 43326474, 108369137, 41357367, 16371544, 39478600, 14213755, 18925372, 16919721, 22428550}
	x := float64(5)
	convey.FocusConvey("TestForecast", t, func() {
		y, err := Forecast(x, knownY, knownX)
		convey.So(err, convey.ShouldBeNil)
		convey.So(math.Round(y), convey.ShouldEqual, float64(26202379))
	})
}
