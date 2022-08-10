package u_context

import (
	"context"
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestContext(t *testing.T) {
	convey.Convey("Test Context", t, func() {
		convey.Convey("Value", func() {
			key := "test_key"
			ctx := Set[string](context.Background(), key, "abc")
			value, ok := Get[string](ctx, key)
			convey.So(ok, convey.ShouldBeTrue)
			convey.So(value, convey.ShouldEqual, "abc")
			fmt.Println(value)
		})

		convey.Convey("Pointer Struct", func() {
			type TestValue struct {
				Name string
				Value int
			}

			key := "test_key"
			ctx := Set[*TestValue](context.Background(), key, &TestValue{
				Name: "abc",
				Value: 64,
			})
			value, ok := Get[*TestValue](ctx, key)
			convey.So(ok, convey.ShouldBeTrue)
			convey.So(value.Name, convey.ShouldEqual, "abc")
			convey.So(value.Value, convey.ShouldEqual, 64)
			fmt.Println(value)
		})
	})
}
