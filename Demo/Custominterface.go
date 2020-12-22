package Demo

import (
	"JdunionSdk_Goalng/JdunionSdk"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type NewJDSDKAPI interface {
	JdunionSdk.JdSdkApi
	NewFunc(s string) *JdunionSdk.CateGoryResult
}

type JdSdkCustom struct {
	JdunionSdk.JdSdk
}

func (J *JdSdkCustom) NewFunc(UriQuery string) *JdunionSdk.CateGoryResult {
	Method := "jd.union.open.category.goods.get"
	J.SetSignJointUrlParam(Method, UriQuery)
	var urls strings.Builder
	urls.WriteString(JdunionSdk.JD_HOST)
	urls.WriteString(J.SignAndUri)
	fmt.Println(urls.String())
	body, _ := JdunionSdk.HttpGet(urls.String())
	result := &JdunionSdk.CategoryResponse{}
	e := json.Unmarshal([]byte(body), &result)
	if e != nil {
		panic(e)
	}
	categoryResult := &JdunionSdk.CateGoryResult{}
	e = json.Unmarshal([]byte(result.JdUnionOpenCategoryGoodsGetResponse.Result), categoryResult)
	if e != nil {
		panic(e)
	}
	return categoryResult
}

type NewJdSdkApi2 struct {
	JdunionSdk.JdSdkApi
}

var _ JdunionSdk.JdSdkApi = &NewJdSdkApi2{}

var JdSdkConfigCustom *JdSdkCustom

func New(AppKeyCustom, AppSecretCustom string) {

	JdunionSdk.JdAppKey = AppKeyCustom
	JdunionSdk.JdAppSecret = AppSecretCustom
	J := &JdSdkCustom{}
	J.RequestParam = JdunionSdk.Param{
		Method:       "",
		App_key:      JdunionSdk.JdAppKey,
		Timestamp:    time.Now().Format("2006-01-02 15:04:05"),
		Format:       "json",
		Access_token: "md5",
		V:            "1.0",
		Param_json:   "",
	}
	JdSdkConfigCustom = J
}

var NewJDSdkCutom NewJDSDKAPI

func init() {
	New(APPKEY, APPSECRET)
	NewJDSdkCutom = JdSdkConfigCustom
}
