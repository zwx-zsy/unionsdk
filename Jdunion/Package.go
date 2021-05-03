package Jdunion

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

//定义接口
type JdSdkApi interface {
	PromotionCommon(query string) *PromotionCommonResult //jd.union.open.promotion.common.get
	OrderRowQuery(query string) *OrderRowQueryResult     //jd.union.open.order.row.query
	OrderBonusQuery(query string) *OrderBonusQueryResult //jd.union.open.order.bonus.query
	ConversionLink(query string) *SubUnionIdResult       //jd.union.open.promotion.bysubunionid.get
	GetJdGoods(query string) *GoodsResult                //jd.union.open.goods.query
	GetCategoryList(query string) *CateGoryResult        //jd.union.open.category.goods.get
	GetGoodsJFen(query string) *JFRestult                //jd.union.open.category.goods.get
	GetOrders(query string) *OrderResult                 //jd.union.open.order.query
	SetSignJointUrlParam(Method string, query string) *JdSdk
}

//var _ JdSdkApi = &JdSdk{}

type Param struct {
	Method       string `json:"method"`
	App_key      string `json:"app_key"`
	Access_token string `json:"access_token"`
	Timestamp    string `json:"timestamp"`
	Format       string `json:"format"`
	V            string `json:"v"`
	Sign_method  string `json:"sign_method"`
	Param_json   string `json:"360buy_param_json"`
}
type JdSdk struct {
	RequestParam Param
	SignAndUri   string
	RequestRul   string
	JdAppKey     string
	AppSecret    string
	JdHost       string
}

func (J *JdSdk) NewContext(AppId, AppSecret string) {
	J.JdHost = "https://api.jd.com/routerjson?"
	location, _ := time.LoadLocation("Asia/Shanghai")
	J.JdAppKey = AppId
	J.AppSecret = AppSecret
	J.RequestParam = Param{
		Method:      "",
		App_key:     J.JdAppKey,
		Timestamp:   time.Now().In(location).Format("2006-01-02 15:04:05"),
		Format:      "json",
		Sign_method: "md5",
		V:           "1.0",
		Param_json:  "",
	}
}

func StructToString(query interface{}) (res string, err error) {
	marshal, err := json.Marshal(query)
	if err != nil {
		return res, err
	}
	return string(marshal), nil
}

func (J *JdSdk) BodyBytes(Method string, query interface{}) []byte {
	toString, err := StructToString(query)
	if err != nil {
		panic(err)
	}
	J.SetSignJointUrlParam(Method, toString)
	urls := strings.Builder{}
	urls.WriteString(J.JdHost)
	urls.WriteString(J.SignAndUri)
	fmt.Println("urls ", urls.String())
	body, _ := HttpGet(urls.String())
	return body
}
