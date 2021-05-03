package Jdunion

import (
	"encoding/json"
)

type Coupon struct {
	BindType     int     `json:"bindType"`
	Discount     float64 `json:"discount"`
	GetEndTime   int64   `json:"getEndTime"`
	GetStartTime int64   `json:"getStartTime"`
	Link         string  `json:"link"`
	PlatformType int     `json:"platformType"`
	Quota        float64 `json:"quota"`
	UseEndTime   int64   `json:"useEndTime"`
	UseStartTime int64   `json:"useStartTime"`
	IsBest       int     `json:"isBest"`
}

type GoodsResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		CategoryInfo struct {
			Cid1 int64 `json:"cid1"`
		} `json:"categoryInfo"`
		BrandCode          string  `json:"brandCode"`
		BrandName          string  `json:"brandName"`
		GoodCommentsShare  float64 `json:"goodCommentsShare"`
		MaterialUrl        string  `json:"materialUrl"`
		Owner              string  `json:"owner"`
		InOrderCount30Days int     `json:"inOrderCount30Days"`
		CouponInfo         struct {
			CouponList []Coupon `json:"couponList"`
		} `json:"couponInfo"`
		Coupon   Coupon `json:"coupon"`
		ShopInfo struct {
			ShopId   int    `json:"shopId"`
			ShopName string `json:"shopName"`
		} `json:"shopInfo"`
		SkuId     int64  `json:"skuId"`
		SkuName   string `json:"skuName"`
		PriceInfo struct {
			Price             float64 `json:"price"`
			LowestPrice       float64 `json:"lowestPrice"`
			LowestPriceType   int     `json:"lowestPriceType"`
			LowestCouponPrice float64 `json:"lowestCouponPrice"`
		} `json:"priceInfo"`
		ImageInfo struct {
			ThumbImg  string `json:"thumbImg"`
			ImageList []struct {
				Url string `json:"url"`
			} `json:"imageList"`
			IsJdSale int `json:"isJdSale"`
		} `json:"imageInfo"`
		Comments       int `json:"comments"`
		CommissionInfo struct {
			Commission             float64 `json:"commission"`
			CommissionShare        float64 `json:"commissionShare"`
			CouponCommission       float64 `json:"couponCommission"`
			PlusCommissionShare    float64 `json:"plusCommissionShare"`
			SuperSubsidies         float64 `json:"superSubsidies"`         //超级补贴
			ShareCommission        float64 `json:"shareCommission"`        //分享佣金
			TallestShareCommission float64 `json:"tallestShareCommission"` //最高分享佣金+超级补贴
			CurrentShareCommission float64 `json:"currentShareCommission"` //当前分享佣金+超级补贴
			IsLock                 float64 `json:"isLock"`
			StartTime              int64   `json:"startTime"`
			EndTime                int64   `json:"endTime"`
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
		Result string `json:"queryResult"`
		Code   string `json:"code"`
	} `json:"jd_union_open_goods_query_responce"`
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
	e := json.Unmarshal(bodyBytes, &result)
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
	HasContent           int     `json:"hasContent,omitempty"`           //1：查询内容商品；其他值过滤掉此入参条件。
	HasBestCoupon        int     `json:"hasBestCoupon,omitempty"`        //1：查询有最优惠券商品；其他值过滤掉此入参条件。
	Pid                  string  `json:"pid,omitempty"`                  //联盟id_应用iD_推广位id
	CouponUrl            string  `json:"couponUrl,omitempty"`            //优惠券链接
}
