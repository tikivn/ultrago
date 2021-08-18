package alert

import (
	"context"
	"fmt"
	"github.com/leekchan/accounting"
	"github.com/slongfield/pyfmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestTelegramImpl(t *testing.T) {
	convey.FocusConvey("TestTelegramImpl", t, func() {
		convey.Convey("TestTelegram_SendMessage", func() {
			SendTeleMessage(context.Background(), "hello world")
		})

		convey.FocusConvey("TestTelegram_FormatMessage", func() {
			ac := accounting.Accounting{Precision: 0, Thousand: ",", Decimal: "."}
			template :=
				`Dear team,
Tiki AFF có program mới như sau: 
				
{programName}: {commissionUpTo:.1f} % ({startDate} - {endDate})

Ngân sách (VNĐ): {totalBudget}

Get Link Program: 
https://affiliate.tiki.com.vn/get-link/program/{programId}
				
Thanks team.`
			message := pyfmt.Must(template,
				map[string]interface{}{
					"programName":    "Program A",
					"commissionUpTo": 20.6,
					"startDate":      "2021-01-01",
					"endDate":        "2021-01-10",
					"totalBudget":    ac.FormatMoney(100000000000000000000.0111),
					"programId":      "111dee9c-9f46-4fc6-8c52-7b772f6f540a",
				})
			fmt.Println(message)
			SendTeleMessage(context.Background(), message)
		})
	})
}
