package Demo

import "JdunionSdk_Goalng/JdunionSdk"

var NewJDSdk JdunionSdk.JdSdkApi

//自定义京东联盟的参数APPKEY、APPSECRET

func init() {
	JdunionSdk.New(APPKEY, APPSECRET)
	NewJDSdk = JdunionSdk.JDSDKConfig
}
