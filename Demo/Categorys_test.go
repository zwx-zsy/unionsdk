package Demo

import (
	"encoding/json"
	"testing"
	"unionsdk/JdunionSdk"
)

func TestExtendJdSdkApi(t *testing.T) {
	var S JdSdkExtend
	S.NewContext(APPKEY, APPSECRET)
	ParamJson := JdunionSdk.PromotionCommonRequest{}
	ParamJson.PromotionCodeReq.SubUnionId = "1u2o3ui123123"
	ParamJson.PromotionCodeReq.SiteId = "4100263590"
	ParamJson.PromotionCodeReq.MaterialId = "https://item.jd.com/69038420388.html"
	marshal, _ := json.Marshal(ParamJson)
	addFunc := S.AddFunc(string(marshal))
	t.Log(addFunc)
}

func TestGoodsQuery(t *testing.T) {
	var J JdSdkExtend
	J.NewContext(APPKEY, APPSECRET)
	ParsmJson := JdunionSdk.GoodsQueryRequest{}
	ParsmJson.GoodsReqDTO.IsCoupon = 1
	ParsmJson.GoodsReqDTO.Keyword = "iphone12"
	ParsmJson.GoodsReqDTO.SkuIds = []int64{}
	ParsmJson.GoodsReqDTO.PageIndex = 1
	ParsmJson.GoodsReqDTO.PageSize = 20
	goods := J.GetJdGoods(ParsmJson)
	t.Log(JdunionSdk.StructToString(goods))
}

func TestCommon(t *testing.T) {
	ParamJson := JdunionSdk.PromotionCommonRequest{}
	ParamJson.PromotionCodeReq.SubUnionId = "1u2o3ui123123"
	ParamJson.PromotionCodeReq.SiteId = "4100263590"
	ParamJson.PromotionCodeReq.MaterialId = "https://item.jd.com/69038420388.html"
	marshal, _ := json.Marshal(ParamJson)
	commonresult := J.PromotionCommon(string(marshal))
	t.Log("commonresult:", commonresult)
}

//查询分类
func TestCateGorys(t *testing.T) {
	ParamJson := "{\"req\":{\"parentId\":0,\"grade\":0}}"
	categorys := J.GetCategoryList(ParamJson)
	t.Log("categorys:", categorys)
}

func TestCustompromotionbysubunionidget(t *testing.T) {
	ParamJson := "{\"req\":{\"parentId\":0,\"grade\":0}}"
	categorys := J.ConversionLink(ParamJson)
	t.Log(categorys)
}

//查询订单
func TestOrder(t *testing.T) {
	ParamStruct := JdunionSdk.OrderParam{}
	ParamStruct.OrderReq.Time = "202012200142"
	ParamStruct.OrderReq.PageNo = 1
	ParamStruct.OrderReq.PageSize = 10
	ParamStruct.OrderReq.Type = 1
	bytes, _ := json.Marshal(ParamStruct)
	t.Log(string(bytes))

	orders := J.GetOrders(string(bytes))
	t.Log(orders)

}

//京粉
func TestJF(t *testing.T) {
	ParamJF := JdunionSdk.ParamJFReq{}
	ParamJF.GoodsReq.PageSize = 20
	ParamJF.GoodsReq.EliteId = 2
	ParamJF.GoodsReq.PageIndex = 1
	bytes, _ := json.Marshal(ParamJF)
	fen := J.GetGoodsJFen(string(bytes))
	t.Log(JdunionSdk.StructToString(fen))
}
