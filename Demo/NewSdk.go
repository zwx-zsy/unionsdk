package Demo

import "unionsdk/JdunionSdk"

var NewJDSdk JdunionSdk.JdSdk

//自定义京东联盟的参数APPKEY、APPSECRET

func init() {
	NewJDSdk.New(APPKEY, APPSECRET)
}
