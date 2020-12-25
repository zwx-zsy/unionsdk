package JdunionSdk

import (
	"encoding/json"
	"fmt"
	"strings"
)

type JdUnionOpenPromotionCommonGetResponse struct {
	JdUnionOpenPromotionCommonGetResponse struct {
		Result string `json:"result"`
		Code   string `json:"code"`
	} `json:"jd_union_open_promotion_common_get_response"`
}

type PromotionCommonResult struct {
	Code int `json:"code"`
	Data struct {
		ClickURL string `json:"clickURL"`
		ShortURL string `json:"shortURL"`
	} `json:"data"`
	Message   string `json:"message"`
	RequestId string `json:"requestId"`
}

type PromotionCommonRequest struct {
	PromotionCodeReq struct {
		MaterialId    string `json:"materialId"`              //推广物料url，例如活动链接、商品链接等；不支持仅传入skuid
		SiteId        string `json:"siteId"`                  //网站ID/APP ID，入口：京东联盟-推广管理-网站管理/APP管理-查看网站ID/APP ID（1、接口禁止使用导购媒体id入参；2、投放链接的网址或应用必须与传入的网站ID/AppID备案一致，否则订单会判“无效-来源与备案网址不符”）
		SubUnionId    string `json:"subUnionId,omitempty"`    //子渠道标识，您可自定义传入字母、数字或下划线，最多支持80个字符，该参数会在订单行查询接口中展示（需申请权限，申请方法请见https://union.jd.com/helpcenter/13246-13247-46301）
		PositionId    int64  `json:"positionId,omitempty"`    //推广位ID
		Ext1          string `json:"ext1,omitempty"`          //系统扩展参数（需申请权限，申请方法请见https://union.jd.com/helpcenter/13246-13247-46301），最多支持40字符，参数会在订单行查询接口中展示
		Pid           string `json:"pid,omitempty"`           //联盟子推客身份标识（不能传入接口调用者自己的pid）
		CouponUrl     string `json:"couponUrl,omitempty"`     //优惠券领取链接，在使用优惠券、商品二合一功能时入参，且materialId须为商品详情页链接
		GiftCouponKey string `json:"giftCouponKey,omitempty"` //礼金批次号
	} `json:"promotionCodeReq"`
}

func (J *JdSdk) GetPromotionCommon(query string) (res *PromotionCommonResult) {
	Method := "jd.union.open.promotion.common.get"
	J.SetSignJointUrlParam(Method, query)
	var urls strings.Builder
	urls.WriteString(JD_HOST)
	urls.WriteString(J.SignAndUri)
	fmt.Println(urls.String())
	body, _ := HttpGet(urls.String())
	fmt.Println(string(body))
	result := &JdUnionOpenPromotionCommonGetResponse{}
	e := json.Unmarshal([]byte(body), &result)
	if e != nil {
		panic(e)
	}
	res = &PromotionCommonResult{}
	e = json.Unmarshal([]byte(result.JdUnionOpenPromotionCommonGetResponse.Result), res)
	if e != nil {
		panic(e)
	}
	return res
	return
}
