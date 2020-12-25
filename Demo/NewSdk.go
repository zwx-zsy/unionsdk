package Demo

import (
	"fmt"
	"unionsdk/JdunionSdk"
)

var NewJDSdk JdunionSdk.JdSdk

//自定义京东联盟的参数APPKEY、APPSECRET

func init() {
	NewJDSdk.NewContext(APPKEY, APPSECRET)
}

type JdSdkExtend struct {
	JdunionSdk.JdSdk
}

//var _ JdunionSdk.JdSdkApi = &JdSdkExtend{}

func (J *JdSdkExtend) AddFunc(query string) *JdunionSdk.PromotionCommonResult {
	CommonResult := J.PromotionCommon(query)
	fmt.Println(CommonResult)
	return CommonResult
}
