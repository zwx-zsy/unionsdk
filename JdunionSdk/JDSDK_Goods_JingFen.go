package JdunionSdk

import (
	"encoding/json"
	"strings"
)

/*
	@name jd.union.open.category.goods.get 商品类目查询
	@des 根据商品的父类目id查询子类目id信息，通常用获取各级类目对应关系，以便将推广商品归类。业务参数parentId、grade都输入0可查询所有一级类目ID，之后再用其作为parentId查询其子类目。
	@EliteId 1-好券商品,2-京粉APP-jingdong.超级大卖场,3-小程序-jingdong.好券商品,4-京粉APP-jingdong.主题聚惠1-jingdong.服装运动,5-京粉APP-jingdong.主题聚惠2-jingdong.精选家电,6-京粉APP-jingdong.主题聚惠3-jingdong.超市,7-京粉APP-jingdong.主题聚惠4-jingdong.居家生活,10-9.9专区,11-品牌好货-jingdong.潮流范儿,12-品牌好货-jingdong.精致生活,13-品牌好货-jingdong.数码先锋,14-品牌好货-jingdong.品质家电,15-京仓配送,16-公众号-jingdong.好券商品,17-公众号-jingdong.9.9,18-公众号-jingdong.京东配送
*/

type ParamJFReq struct {
	GoodsReq GoodsReq `json:"goodsReq"`
}

type GoodsReq struct {
	EliteId   int    `json:"eliteId"`             //1-好券商品,2-京粉APP-jingdong.超级大卖场,3-小程序-jingdong.好券商品,4-京粉APP-jingdong.主题聚惠1-jingdong.服装运动,5-京粉APP-jingdong.主题聚惠2-jingdong.精选家电,6-京粉APP-jingdong.主题聚惠3-jingdong.超市,7-京粉APP-jingdong.主题聚惠4-jingdong.居家生活,10-9.9专区,11-品牌好货-jingdong.潮流范儿,12-品牌好货-jingdong.精致生活,13-品牌好货-jingdong.数码先锋,14-品牌好货-jingdong.品质家电,15-京仓配送,16-公众号-jingdong.好券商品,17-公众号-jingdong.9.9,18-公众号-jingdong.京东配送
	PageIndex int    `json:"pageIndex,omitempty"` //页码
	PageSize  int    `json:"pageSize,omitempty"`  //每页数量
	SortName  string `json:"sortName,omitempty"`  //排序字段(price：单价, commissionShare：佣金比例, commission：佣金， inOrderCount30DaysSku：sku维度30天引单量，comments：评论数，goodComments：好评数)
	Sort      string `json:"Sort,omitempty"`      //asc,desc升降序,默认降序
}

type JdUnionOpenGoodsJingfenQueryResponse struct {
	JdUnionOpenGoodsJingfenQueryResponse struct {
		Result string `json:"result"`
		Code   string `json:"code"`
	} `json:"jd_union_open_goods_jingfen_query_response"`
}

type JFRestult struct {
	Code int `json:"code"`
	Data []struct {
		CategoryInfo struct {
			Cid1     int    `json:"cid1"`
			Cid1Name string `json:"cid1Name"`
			Cid2     int    `json:"cid2"`
			Cid2Name string `json:"cid2Name"`
			Cid3     int    `json:"cid3"`
			Cid3Name string `json:"cid3Name"`
		} `json:"categoryInfo"`
		Comments       int `json:"comments"`
		CommissionInfo struct {
			Commission      float64 `json:"commission"`
			CommissionShare float64 `json:"commissionShare"`
		} `json:"commissionInfo"`
		CouponInfo struct {
			CouponList []struct {
				BindType     int     `json:"bindType"`
				Discount     float64 `json:"discount"`
				GetEndTime   int64   `json:"getEndTime"`
				GetStartTime int64   `json:"getStartTime"`
				Link         string  `json:"link"`
				PlatformType int     `json:"platformType"`
				Quota        float64 `json:"quota"`
				UseEndTime   int64   `json:"useEndTime"`
				UseStartTime int64   `json:"useStartTime"`
			}
		} `json:"couponInfo"`
		GoodCommentsShare float64 `json:"goodCommentsShare"`
		ImageInfo         struct {
			ImageList []struct {
				Url string `json:"url"`
			}
		} `json:"imageinfo"`
		InOrderCount30Days int64  `json:"inOrderCount30Days"`
		MaterialUrl        string `json:"materialUrl"`
		PriceInfo          struct {
			Price float64 `json:"price"`
		} `json:"priceInfo"`
		ShopInfo struct {
			ShopName string `json:"shopName"`
			ShopId   int64  `json:"shopId"`
		} `json:"shopinfo"`
		SkuId      int64  `json:"skuId"`
		SkuName    string `json:"skuName"`
		IsHot      int    `json:"isHot"`
		Spuid      int64  `json:"spuid"`
		BrandCode  string `json:"brandCode"`
		BrandName  string `json:"brandName"`
		Owner      string `json:"owner"`
		PinGouInfo struct {
			PingouPrice   float64 `json:"pingouPrice"`
			PingouTmCount int64   `json:"pingouTmCount"`
			PingouUrl     string  `json:"pingouUrl"`
		} `json:"pinGouInfo"`
		ResourceInfo struct {
			EliteId   int    `json:"eliteId"`
			EliteName string `json:"eliteName"`
		} `json:"resourceInfo"`
		InOrderCount30DaysSku int64 `json:"inOrderCount30DaysSku"`
	} `json:"data"`
	TotalCount int64  `json:"totalCount"`
	Message    string `json:"message"`
	RequestId  string `json:"requestId"`
}

func (J *JdSdk) GetGoodsJFen(param string) (goodresult *JFRestult) {
	Method := "jd.union.open.goods.jingfen.query"
	J.SetSignJointUrlParam(Method, param)
	var urls strings.Builder
	urls.WriteString(JD_HOST)
	urls.WriteString(J.SignAndUri)
	body, _ := HttpGet(urls.String())
	result := &JdUnionOpenGoodsJingfenQueryResponse{}
	e := json.Unmarshal([]byte(body), &result)
	if e != nil {
		panic(e)
	}
	goodresult = &JFRestult{}
	e = json.Unmarshal([]byte(result.JdUnionOpenGoodsJingfenQueryResponse.Result), goodresult)
	if e != nil {
		panic(e)
	}
	return goodresult
}
