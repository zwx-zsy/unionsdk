package JdunionSdk

import (
	"encoding/json"
)

type GoodsResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		BrandCode         string  `json:"brandCode"`
		BrandName         string  `json:"brandName"`
		GoodCommentsShare float64 `json:"goodCommentsShare"`
		CouponInfo        struct {
			CouponList []struct {
				Discount float64 `json:"discount"`
				Link     string  `json:"link"`
				IsBest   int     `json:"isBest"`
			} `json:"couponList"`
		} `json:"couponInfo"`
		ShopInfo struct {
			ShopId   int    `json:"shopId"`
			ShopName string `json:"shopName"`
		} `json:"shopInfo"`
		SkuId     int64  `json:"skuId"`
		SkuName   string `json:"skuName"`
		PriceInfo struct {
			Price       float64 `json:"price"`
			LowestPrice float64 `json:"lowestprice"`
		} `json:"priceInfo"`
		ImageInfo struct {
			ImageList []struct {
				Url string `json:"url"`
			} `json:"imageList"`
			IsJdSale int `json:"isJdSale"`
		} `json:"imageInfo"`
		Comments       int `json:"comments"`
		CommissionInfo struct {
			Commission      float64 `json:"commission"`
			CommissionShare float64 `json:"commissionshare"`
		} `json:"commissionInfo"`
		DocumentInfo struct {
			Document string `json:"document"`
			Discount string `json:"discount"`
		} `json:"documentInfo"`
	} `json:"data"`
	TotalCount int    `json:"totalCount"`
	RequestId  string `json:"requestId"`
}

//获取商品列表
type JdUnionOpenGoodsQueryResponse struct {
	JdUnionOpenGoodsQueryResponse struct {
		Result string `json:"result"`
		Code   string `json:"code"`
	} `json:"jd_union_open_goods_query_response"`
}

//请求参数
type GoodsQueryRequest struct {
	GoodsReqDTO GoodsReqDTO `json:"goodsReqDTO"`
}

//关键词字商品查询
func (J *JdSdk) GetJdGoods(query interface{}) (GoodResult *GoodsResult) {
	Method := "jd.union.open.goods.query"
	bodyBytes := J.BodyBytes(Method, query)
	result := &JdUnionOpenGoodsQueryResponse{}
	e := json.Unmarshal([]byte(bodyBytes), &result)
	if e != nil {
		panic(e)
	}
	GoodResult = &GoodsResult{}
	e = json.Unmarshal([]byte(result.JdUnionOpenGoodsQueryResponse.Result), GoodResult)
	if e != nil {
		panic(e)
	}
	return GoodResult
}

type GoodsReqDTO struct {
	Cid1                 int64   `json:"cid1,omitempty"`                 //一级类目id
	Cid2                 int64   `json:"cid2,omitempty"`                 //二级类目id
	Cid3                 int64   `json:"cid3,omitempty"`                 //三级类目id
	PageIndex            int     `json:"pageIndex,omitempty"`            //页码
	PageSize             int     `json:"pageSize,omitempty"`             //每页数量，单页数最大30，默认20
	SkuIds               []int64 `json:"skuIds,omitempty"`               //skuid 集合(一次最多支持查询100个sku)，数组类型开发时记得加[]
	Keyword              string  `json:"keyword,omitempty"`              //关键词，字数同京东商品名称一致，目前未限制
	Pricefrom            float64 `json:"pricefrom,omitempty"`            //商品价格下限
	Priceto              float64 `json:"priceto,omitempty"`              //商品价格上限
	CommissionShareStart int     `json:"commissionShareStart,omitempty"` //佣金比例区间开始
	CommissionShareEnd   int     `json:"commissionShareEnd,omitempty"`   // 佣金比例区间结束
	Owner                string  `json:"owner,omitempty"`                //商品类型：自营[g]，POP[p]
	SortName             string  `json:"sortName,omitempty"`             //排序字段(price：单价, commissionShare：佣金比例, commission：佣金， inOrderCount30Days：30天引单量， inOrderComm30Days：30天支出佣金)
	Sort                 string  `json:"sort,omitempty"`                 //asc,desc升降序,默认降序
	IsCoupon             int     `json:"isCoupon,omitempty"`             //是否是优惠券商品，1：有优惠券，0：无优惠券
	IsPG                 int     `json:"isPG,omitempty"`                 //是否是拼购商品，1：拼购商品，0：非拼购商品
	PingouPriceStart     float64 `json:"pingouPriceStart,omitempty"`     //拼购价格区间开始
	PingouPriceEnd       float64 `json:"pingouPriceEnd,omitempty"`       //拼购价格区间结束
	IsHot                int     `json:"isHot,omitempty"`                //是否是爆款，1：爆款商品，0：非爆款商品
	BrandCode            string  `json:"brandCode,omitempty"`            //品牌code
	ShopId               int     `json:"shopId,omitempty"`               //店铺Id
}
