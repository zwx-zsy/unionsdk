package Demo

import (
	"encoding/json"
	"fmt"
	"testing"
	"unionsdk/JdunionSdk"
)

func TestCutomInterfaceGoods(t *testing.T) {
	//var Newsapi NewJdSdkApi2
	var S JdSdkExtend
	S.NewContext(APPKEY, APPSECRET)
	//ParamJson := "{\"req\":{\"parentId\":0,\"grade\":0}}"
	ParamJson := JdunionSdk.PromotionCommonRequest{}
	ParamJson.PromotionCodeReq.SubUnionId = "1u2o3ui123123"
	ParamJson.PromotionCodeReq.SiteId = "4100263590"
	ParamJson.PromotionCodeReq.MaterialId = "https://item.jd.com/69038420388.html"
	marshal, _ := json.Marshal(ParamJson)
	fmt.Println(string(marshal))
	addFunc := S.AddFunc(string(marshal))
	fmt.Println(addFunc)
}

func TestGoodsQuery(t *testing.T) {
	var J JdSdkExtend
	J.NewContext(APPKEY, APPSECRET)
	//addFunc := J.AddFunc("s")
	//fmt.Println(addFunc)
	ParsmJson := JdunionSdk.GoodsQueryRequest{}
	ParsmJson.GoodsReqDTO.IsCoupon = 1
	ParsmJson.GoodsReqDTO.Keyword = "iphone12"
	ParsmJson.GoodsReqDTO.SkuIds = []int64{}
	ParsmJson.GoodsReqDTO.PageIndex = 1
	ParsmJson.GoodsReqDTO.PageSize = 20
	goods := J.GetJdGoods(ParsmJson)
	fmt.Println(JdunionSdk.StructToString(goods))
}

func TestCommon(t *testing.T) {
	ParamJson := JdunionSdk.PromotionCommonRequest{}
	ParamJson.PromotionCodeReq.SubUnionId = "1u2o3ui123123"
	ParamJson.PromotionCodeReq.SiteId = "4100263590"
	ParamJson.PromotionCodeReq.MaterialId = "https://item.jd.com/69038420388.html"
	marshal, _ := json.Marshal(ParamJson)
	fmt.Println(string(marshal))
	//categorys2 := NewJDSdk.GetPromotionCommon(string(marshal))
	//t.Log(categorys2)
	//fmt.Println(categorys2)
	var J JdSdkExtend
	J.NewContext(APPKEY, APPSECRET)
	addFunc := J.AddFunc(string(marshal))
	fmt.Println(addFunc)
}

//查询分类
func TestCateGorys(t *testing.T) {
	ParamJson := "{\"req\":{\"parentId\":0,\"grade\":0}}"
	categorys := NewJDSdk.GetCategoryList(ParamJson)
	t.Log(categorys)
}

func TestCustompromotionbysubunionidget(t *testing.T) {
	ParamJson := "{\"req\":{\"parentId\":0,\"grade\":0}}"
	categorys := NewJDSdk.ConversionLink(ParamJson)
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
	fmt.Println(string(bytes))

	orders := NewJDSdk.GetOrders(string(bytes))
	fmt.Println(orders)

}

//京粉
func TestJF(t *testing.T) {
	ParamJF := JdunionSdk.ParamJFReq{}
	ParamJF.GoodsReq.PageSize = 20
	ParamJF.GoodsReq.EliteId = 2
	ParamJF.GoodsReq.PageIndex = 1
	bytes, _ := json.Marshal(ParamJF)
	fen := NewJDSdk.GetGoodsJFen(string(bytes))
	t.Log(fen)
}
