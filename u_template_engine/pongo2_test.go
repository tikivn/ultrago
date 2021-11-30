package u_template_engine

import (
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestParse(t *testing.T) {
	text1 := "[Tiki AFF] - Đối soát hoa hồng cơ bản kì [Tháng {{month}}/{{year}}] đã được tạo"
	text2 := `
<!DOCTYPE html><html><head> <meta charset="utf-8" /> <meta http-equiv="x-ua-compatible" content="ie=edge" /> <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no" /> <style data-href="/styles.96f838a0b45813725758.css" data-identity="gatsby-global-css"> body { box-sizing: border-box !important; font-family: sans-serif; font-size: 14px; font-weight: 400; line-height: 1.5; margin: 0; } </style> <meta name="generator" content="Gatsby 3.14.2" /> <style data-styled="" data-styled-version="5.3.1"> .cykBeC { width: 100vw; height: 100vh; background-color: #eeeeee; } /*!sc*/ .emEDuF { margin-left: auto; margin-right: auto; width: 100%; background-color: white; } /*!sc*/ @media screen and (min-width: 40em) { .emEDuF { width: 500px; } } /*!sc*/ .ithjoq { padding: 24px; } /*!sc*/ .eufwBf { display: -webkit-box; display: -webkit-flex; display: -ms-flexbox; display: flex; -webkit-align-items: center; -webkit-box-align: center; -ms-flex-align: center; align-items: center; -webkit-box-pack: justify; -webkit-justify-content: space-between; -ms-flex-pack: justify; justify-content: space-between; -webkit-flex-direction: row; -ms-flex-direction: row; flex-direction: row; -webkit-flex-wrap: wrap; -ms-flex-wrap: wrap; flex-wrap: wrap; } /*!sc*/ .liHqUT { display: -webkit-box; display: -webkit-flex; display: -ms-flexbox; display: flex; -webkit-flex-shrink: 1; -ms-flex-negative: 1; flex-shrink: 1; } /*!sc*/ .bBoIhA { display: -webkit-box; display: -webkit-flex; display: -ms-flexbox; display: flex; -webkit-align-items: center; -webkit-box-align: center; -ms-flex-align: center; align-items: center; } /*!sc*/ .gEOOgg { margin-top: 12px; margin-right: -24px; margin-left: -24px; height: 3px; background-color: #00b7e8; } /*!sc*/ .CzDCa { padding-top: 8px; } /*!sc*/ .dAhLJu { padding: 16px; color: #999999; font-size: 12px; text-align: center; } /*!sc*/ data-styled.g1[id="Box-sc-1sgkirf-0"] { content: "cykBeC,emEDuF,ithjoq,eufwBf,liHqUT,bBoIhA,gEOOgg,CzDCa,dAhLJu,"; } /*!sc*/ </style></head><body> <div id="___gatsby"> <div style="outline: none" tabindex="-1" id="gatsby-focus-wrapper"> <div width="100vw" height="100vh" class="Box-sc-1sgkirf-0 cykBeC"> <div width="100%,500" class="Box-sc-1sgkirf-0 emEDuF"> <div class="Box-sc-1sgkirf-0 ithjoq"> <div display="flex" class="Box-sc-1sgkirf-0 eufwBf"> <div display="flex" class="Box-sc-1sgkirf-0 liHqUT"> <a target="_blank" href="https://tiki.vn"><img style="height: 40px; margin-bottom: 8px" src="https://ci5.googleusercontent.com/proxy/LKEDgkrP_XHtlP-ulbE5_RJAh1ivbpYjzxwQRWGizJxTGcxJl0LDSbA8r53ij7T9ZNcwUc-TJOEK=s0-d-e1-ft#http://tikicdn.com/assets/img/logo.png" /></a> </div> <div display="flex" class="Box-sc-1sgkirf-0 bBoIhA"> <a target="_blank" style=" text-decoration: none; margin-right: 8px; margin-bottom: 8px; " href="https://apps.apple.com/vn/app/tiki-shopping-fast-shipping/id958100553"><img alt="Tải ứng dụng iOS" src="https://ci4.googleusercontent.com/proxy/lrppks7xzE6euW2045YpDJelaAHNnbav0B5_ZEIxDaODUkTKJWrN6G_WczSY9Gh_bKNFWt3Pr6HvFib3aMdGgduEsqqoDZ4=s0-d-e1-ft#http://tikicdn.com/media/custom/app_store_ios_2x.png" style="height: 40px" /></a><a target="_blank" href="https://play.google.com/store/apps/details?id=vn.tiki.app.tikiandroid&amp;hl=en&amp;gl=US" style="text-decoration: none; margin-bottom: 8px"><img alt="Tải ứng dụng Android" src="https://ci5.googleusercontent.com/proxy/uY_osf13Vr_P4EHX7czu7re4LKf4PuBbvuSSTLV2NaP7JrNIOPGEOgv5IIBbzg_0Pwf6Uk92Ieu6jWKpEIzC8Mq-TEM=s0-d-e1-ft#http://tikicdn.com/media/custom/play_store_2x.png" style="height: 40px" /></a> </div> </div> <div height="3px" class="Box-sc-1sgkirf-0 gEOOgg"></div> <!--START_CONTENT--> <div class="Box-sc-1sgkirf-0 CzDCa"> {% autoescape off %} <p>{{ greeting }}</p> {% for block in block_stmt %} <p> {{ block.stmt1 }} {% if block.stmt2 %} <br>{{ block.stmt2 }} {% endif %} {% if block.has_link %} <a href="{{ target_url }}" target="_blank" style="font-weight: bold; color: #00b7e8">đây</a> {% endif %} </p> {% endfor %} <p> </p> <p style="margin-bottom: 0"> Trân trọng,<br> Đội ngũ Tiki Affiliate. </p> {% endautoescape %} </div> <!--END_CONTENT--> </div> </div> <div color="#999999" font-size="12px" class="Box-sc-1sgkirf-0 dAhLJu"> Công ty Cổ phần Tiki (52 Út Tịch, P.4, Q.Tân Bình, Tp.HCM) </div> </div> </div> <div id="gatsby-announcer" style=" position: absolute; top: 0; width: 1px; height: 1px; padding: 0; overflow: hidden; clip: rect(0, 0, 0, 0); white-space: nowrap; border: 0; " aria-live="assertive" aria-atomic="true"></div> </div></body></html>
`
	text3 := "Raw text: {{rawText}}"
	text4 := "Nested array access: {{list.0.name}}"

	convey.Convey("TestParse", t, func() {
		res1, err1 := Parse(text1, map[string]interface{}{
			"month": 9,
			"year":  2021,
		})
		convey.So(err1, convey.ShouldBeNil)
		convey.So(res1, convey.ShouldEqual, "[Tiki AFF] - Đối soát hoa hồng cơ bản kì [Tháng 9/2021] đã được tạo")

		res2, err2 := Parse(text2, map[string]interface{}{
			"greeting": "Kính chào quý đối tác {{username}},",
			"block_stmt": []map[string]interface{}{
				{
					"stmt1":    "Đối soát hoa hồng kì <b>[Tháng {{month}}/{{year}}]</b> đã được tạo.",
					"stmt2":    "Vui lòng truy cập đường dẫn sau đây để biết thêm chi tiết:",
					"has_link": "yes",
				},
			},
			"username":   "Nguyễn Văn A",
			"month":      9,
			"year":       2021,
			"target_url": "https://example.com/",
		})
		convey.So(err2, convey.ShouldBeNil)
		fmt.Println(res2)

		res3, err3 := Parse(text3, map[string]interface{}{
			"rawText":   "Final text: {{finalText}}",
			"finalText": "Hihihi",
		})
		convey.So(err3, convey.ShouldBeNil)
		convey.So(res3, convey.ShouldEqual, "Raw text: Final text: Hihihi")

		res4, err4 := Parse(text4, map[string]interface{}{
			"list": []map[string]string{
				{
					"name": "abc",
				},
			},
		})
		convey.So(err4, convey.ShouldBeNil)
		convey.So(res4, convey.ShouldEqual, "Nested array access: abc")
	})
}
