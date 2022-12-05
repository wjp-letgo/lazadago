package config

import(
	"testing"
	"github.com/wjpxxx/letgo/lib"
	"fmt"
)
func TestConfig(t *testing.T){
	var t2 int64=1629265567128
	xml:="<Request><Product><PrimaryCategory>6614</PrimaryCategory><SPUId/><AssociatedSku/><Images><Image>https://my-live-02.slatic.net/p/765888ef9ec9e81106f451134c94048f.jpg</Image><Image>https://my-live-02.slatic.net/p/9eca31edef9f05f7e42f0f19e4d412a3.jpg</Image></Images><Attributes><name>api create product test sample</name><short_description>This is a nice product</short_description><brand>Remark</brand><model>asdf</model><kid_years>Kids (6-10yrs)</kid_years><video>12345 (fill with the video id of the previously uploaded video) optional</video><delivery_option_sof>Yes</delivery_option_sof></Attributes><Skus><Sku><SellerSku>api-create-test-1</SellerSku><color_family>Green</color_family><size>40</size><quantity>1</quantity><price>388.50</price><package_length>11</package_length><package_height>22</package_height><package_weight>33</package_weight><package_width>44</package_width><package_content>this is what's in the box</package_content><Images><Image>http://sg.s.alibaba.lzd.co/original/59046bec4d53e74f8ad38d19399205e6.jpg</Image><Image>http://sg.s.alibaba.lzd.co/original/179715d3de39a1918b19eec3279dd482.jpg</Image></Images></Sku></Skus></Product></Request>"
	fmt.Println(Sign("EbMYJasQyI0zJQA7IDvbScsq3tzaxmMW","/product/create",lib.InRow{
		"app_key":113575,
		"sign_method":"sha256",
		"access_token":"werwerwerwerwerwerwr",
		"timestamp":t2,
		"payload":xml,
		"@image":[]byte{},
	}))
}