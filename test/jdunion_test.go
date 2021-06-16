package jdunion_test

import (
	"testing"
	Jd "unionsdk/Jdunion"

	"github.com/golang/glog"
)

const (
	AppId     string = ""
	AppSecret string = ""
)

var Jds Jd.JdSdk

//京粉精选商品查询接口
func TestGetGoodsJFen(t *testing.T) {
	Jds.NewContext(AppId, AppSecret)
	var req Jd.ParamJFReq
	req.GoodsReq.EliteId = 10
	res := Jds.GetGoodsJFen(req)
	glog.Infoln(res)
}

//转链
func TestBySubUnionId(t *testing.T) {
	Jds.NewContext(AppId, AppSecret)
	var req Jd.BySubUnionIdRequest
	req.PromotionCodeReq.MaterialId = "https://item.jd.com/100016889372.html"
	req.PromotionCodeReq.SubUnionId = "18856988766"
	res := Jds.ConversionLink(req)
	glog.Infoln(res)
}

//获取订单行数据
func TestOrderRowQuery(t *testing.T) {
	Jds.NewContext(AppId, AppSecret)
	var req Jd.OrderRowQueryRequest
	req.OrderReq.Fields = "goodsInfo"
	req.OrderReq.StartTime = "2021-05-03 21:40:00"
	req.OrderReq.EndTime = "2021-05-03 21:52:00"
	req.OrderReq.Type = 3
	req.OrderReq.PageIndex = 1
	req.OrderReq.PageSize = 200
	rowQuery := Jds.OrderRowQuery(req)
	glog.Infoln(rowQuery)
}

//jd.union.open.goods.query
func TestGoodsQuery(t *testing.T) {
	Jds.NewContext(AppId, AppSecret)
	var req Jd.GoodsQueryRequest
	req.GoodsReqDTO.PageIndex = 1
	req.GoodsReqDTO.PageSize = 200
	req.GoodsReqDTO.Keyword = "iphone"
	goods := Jds.GetJdGoods(req)
	glog.Infoln(goods)
}

//商品类目
func TestCategory(t *testing.T) {
	Jds.NewContext(AppId, AppSecret)
	var req Jd.CategoryRequest
	req.Req.ParentId = 0
	categoryList := Jds.GetCategoryList(req)
	glog.Infoln(categoryList)
}
