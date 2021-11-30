package u_miscellaneous

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/smartystreets/goconvey/convey"
)

func TestUUID2UInt(t *testing.T) {
	convey.Convey("TestUUID2UInt", t, func() {
		existed := make(map[uint64]bool, 0)
		for i := 0; i < 1000000; i++ {
			uid, err := UUID2UInt(uuid.New().String())
			convey.So(err, convey.ShouldBeNil)
			convey.So(len(fmt.Sprintf("%d", uid)), convey.ShouldBeLessThanOrEqualTo, 20)
			_, ok := existed[uid]
			convey.So(ok, convey.ShouldBeFalse)
			existed[uid] = true
		}
	})
}
