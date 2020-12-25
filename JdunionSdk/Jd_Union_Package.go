package JdunionSdk

import "time"

const (
	JD_HOST            string = "https://router.jd.com/api?"
	CUSTOMREQMETHODGET string = "GET"
)

//定义接口
type JdSdkApi interface {
	//GoodsMaterialQuery(query string) *GoodsMaterialQueryResult //jd.union.open.goods.material.query
	//OrderBonusQuery(query string) *OrderBonusQueryResult       //jd.union.open.order.bonus.query
	//OrderRowQuery(query string) *OrderRowQueryResult           //jd.union.open.order.row.query
	//GetPromotionCommon(query string) *PromotionCommonResult    //jd.union.open.promotion.common.get
	//ConversionLink(query string) *SubUnionIdResult             //jd.union.open.promotion.bysubunionid.get
	//GetJdGoods(query string) *GoodsResult                      //jd.union.open.goods.query
	//GetCategoryList(query string) *CateGoryResult              //jd.union.open.category.goods.get
	//GetGoodsJFen(query string) *JFRestult                      //jd.union.open.goods.jingfen.query
	//GetOrders(query string) *OrderResult                       //jd.union.open.order.query
	//SetSignJointUrlParam(Method string, query string) *JdSdk   //携带请求参数签名
}

type Param struct {
	Method       string `json:"method"`
	App_key      string `json:"app_key"`
	Access_token string `json:"access_token"`
	Timestamp    string `json:"timestamp"`
	Format       string `json:"format"`
	V            string `json:"v"`
	Sign_method  string `json:"sign_method"`
	Param_json   string `json:"param_json"`
}

type JdSdk struct {
	RequestParam Param
	SignAndUri   string
	RequestRul   string
}

var _ JdSdkApi = &JdSdk{}

var JdAppKey, JdAppSecret string

func (J *JdSdk) New(AppId, AppSecret string) {
	JdAppKey = AppId
	JdAppSecret = AppSecret
	J.RequestParam = Param{
		Method:      "",
		App_key:     JdAppKey,
		Timestamp:   time.Now().Format("2006-01-02 15:04:05"),
		Format:      "json",
		Sign_method: "md5",
		V:           "1.0",
		Param_json:  "",
	}
}
