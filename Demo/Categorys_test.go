package Demo

import (
	"encoding/json"
	"testing"
	"unionsdk/JdunionSdk"
)

func TestExtendJdSdkApi(t *testing.T) {
	var S JdSdkExtend
	S.NewContext(APPKEY, APPSECRET)
	PromotionCommonParam := JdunionSdk.PromotionCommonRequest{}
	PromotionCommonParam.PromotionCodeReq.SubUnionId = "1u2o3ui123123"
	PromotionCommonParam.PromotionCodeReq.SiteId = "4100263590"
	PromotionCommonParam.PromotionCodeReq.MaterialId = "https://item.jd.com/69038420388.html"
	addFunc := S.AddFunc(PromotionCommonParam)
	t.Log(addFunc)
}

func TestGoodsQuery(t *testing.T) {
	ParamGoodsQuery := JdunionSdk.GoodsQueryRequest{}
	ParamGoodsQuery.GoodsReqDTO.IsCoupon = 1
	ParamGoodsQuery.GoodsReqDTO.Keyword = "iphone12"
	ParamGoodsQuery.GoodsReqDTO.SkuIds = []int64{}
	ParamGoodsQuery.GoodsReqDTO.PageIndex = 1
	ParamGoodsQuery.GoodsReqDTO.PageSize = 20
	goods := J.GetJdGoods(ParamGoodsQuery)
	t.Log(JdunionSdk.StructToString(goods))
}

func TestPromotionCommon(t *testing.T) {
	PromotionCommonParam := JdunionSdk.PromotionCommonRequest{}
	PromotionCommonParam.PromotionCodeReq.SubUnionId = "1u2o3ui123123"
	PromotionCommonParam.PromotionCodeReq.SiteId = "4100263590"
	PromotionCommonParam.PromotionCodeReq.MaterialId = "https://item.jd.com/69038420388.html"
	marshal, _ := json.Marshal(PromotionCommonParam)
	promotionCommon := J.PromotionCommon(string(marshal))
	t.Log("promotionCommon:", promotionCommon)
}

//查询分类
func TestCateGory(t *testing.T) {
	ParamCateGory := JdunionSdk.CategoryRequest{}
	categoryResult := J.GetCategoryList(ParamCateGory)
	t.Log("categoryResult:", categoryResult)
}

func TestPromotionBySubUnionIdGet(t *testing.T) {
	ParamBySubUnionId := struct {
		PromotionCodeReq JdunionSdk.PromotionCodeReq `json:"promotionCodeReq"`
	}{}
	ParamBySubUnionId.PromotionCodeReq.MaterialId = "item.jd.com/10023730856633.html"
	ParamBySubUnionId.PromotionCodeReq.CouponUrl = "https://coupon.m.jd.com/coupons/show.action?linkKey=AAROH_xIpeffAs_-naABEFoePOfM-rS0222iXTc6vdhzwurvzgakr3brWUXJrzQY0L97qDKYHXsZ-ty4Da1p_c2degL7Mg"
	subUnionIdResult := J.ConversionLink(ParamBySubUnionId)
	t.Log(subUnionIdResult)
}

//查询订单
func TestOrder(t *testing.T) {
	ParamOrder := JdunionSdk.OrderParam{}
	ParamOrder.OrderReq.Time = "202012210142"
	ParamOrder.OrderReq.PageNo = 1
	ParamOrder.OrderReq.PageSize = 10
	ParamOrder.OrderReq.Type = 3
	orders := J.GetOrders(ParamOrder)
	t.Log(orders)

}

//京粉
func TestJF(t *testing.T) {
	ParamJF := JdunionSdk.ParamJFReq{}
	ParamJF.GoodsReq.PageSize = 20
	ParamJF.GoodsReq.EliteId = 2
	ParamJF.GoodsReq.PageIndex = 1
	fen := J.GetGoodsJFen(ParamJF)
	t.Log(JdunionSdk.StructToString(fen))
}
