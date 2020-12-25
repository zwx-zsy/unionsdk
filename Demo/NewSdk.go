package Demo

import (
	"fmt"
	"unionsdk/JdunionSdk"
)

var J JdunionSdk.JdSdk

//自定义京东联盟的参数APPKEY、APPSECRET

func init() {
	J.NewContext(APPKEY, APPSECRET)
}

type JdSdkExtend struct {
	JdunionSdk.JdSdk
}

func (J *JdSdkExtend) AddFunc(query interface{}) *JdunionSdk.PromotionCommonResult {
	CommonResult := J.PromotionCommon(query)
	fmt.Println(CommonResult)
	return CommonResult
}
