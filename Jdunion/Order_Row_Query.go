package Jdunion

import (
	"encoding/json"
)

type OrderRowQueryRequest struct {
	OrderReq struct {
		PageIndex    int    `json:"pageIndex"`              //页码，返回第几页结果
		PageSize     int    `json:"pageSize"`               //每页包含条数，上限为500
		Type         int    `json:"type"`                   //订单时间查询类型(1：下单时间，2：完成时间（购买用户确认收货时间），3：更新时间
		StartTime    string `json:"startTime"`              //开始时间 格式yyyy-MM-dd HH:mm:ss，与endTime间隔不超过1小时
		EndTime      string `json:"endTime"`                //结束时间 格式yyyy-MM-dd HH:mm:ss，与startTime间隔不超过1小时
		ChildUnionId int64  `json:"childUnionId,omitempty"` //子站长ID（需要联系运营开通PID账户权限才能拿到数据），childUnionId和key不能同时传入
		Key          string `json:"key,omitempty"`          //xxxxxxab4d5247fdac36ddbaf0xxxxxx 工具商传入推客的授权key，可帮助该推客查询订单，注意不可和childUnionid同时传入。（需联系运营开通工具商权限才能拿到数据）
		Fields       string `json:"fields,omitempty"`       //例如:goodsInfo .支持出参数据筛选，逗号','分隔，目前可用：goodsInfo（商品信息）,categoryInfo(类目信息）
	} `json:"orderReq"`
}

type JdUnionOpenOrderRowQueryResponse struct {
	JdUnionOpenOrderRowQueryResponse struct {
		Result string `json:"queryResult"`
		Code   string `json:"code"`
	} `json:"jd_union_open_order_row_query_responce"`
}

type OrderRowQueryResult struct {
	Code      int    `json:"code"`    //返回码
	HasMore   bool   `json:"hasMore"` //是否还有更多,true：还有数据；false:已查询完毕，没有数据
	Message   string `json:"message"` //返回消息
	RequestId string `json:"requestId"`
	Data      []struct {
		Id                  string  `json:"id"`                  //标记唯一订单行：订单+sku维度的唯一标识
		OrderId             int64   `json:"orderId"`             //订单ID
		ParentId            int64   `json:"parentId"`            //父单的订单ID，仅当发生订单拆分时返回， 0：未拆分，有值则表示此订单为子订单
		OrderTime           string  `json:"orderTime"`           //下单时间(时间戳，毫秒)
		FinishTime          string  `json:"finishTime"`          //订单完成时间(时间戳，毫秒)
		ModifyTime          string  `json:"modifyTime"`          //下单时间(时间戳，毫秒)
		OrderEmt            int     `json:"orderEmt"`            //下单设备(1:PC,2:无线)
		Plus                int     `json:"plus"`                //下单用户是否为PLUS会员 0：否，1：是
		UnionId             int64   `json:"unionId"`             //推客ID
		SkuId               int64   `json:"skuId"`               //商品ID
		SkuName             string  `json:"skuName"`             //商品名称
		SkuNum              int     `json:"skuNum"`              //商品数量
		SkuReturnNum        int     `json:"skuReturnNum"`        //商品已退货数量
		SkuFrozenNum        int     `json:"skuFrozenNum"`        //商品售后中数量
		Price               float64 `json:"price"`               //商品单价
		CommissionRate      float64 `json:"commissionRate"`      //佣金比例(投放的广告主计划比例)
		SubSideRate         float64 `json:"subSideRate"`         //分成比例（单位：%）
		SubsidyRate         float64 `json:"subsidyRate"`         //补贴比例（单位：%）
		FinalRate           float64 `json:"finalRate"`           //最终分佣比例（单位：%）=分成比例+补贴比例
		EstimateCosPrice    float64 `json:"estimateCosPrice"`    //预估计佣金额：由订单的实付金额拆分至每个商品的预估计佣金额，不包括运费，以及京券、东券、E卡、余额等虚拟资产支付的金额。该字段仅为预估值，实际佣金以actualCosPrice为准进行计算
		EstimateFee         float64 `json:"estimateFee"`         //推客的预估佣金（预估计佣金额*佣金比例*最终比例），如订单完成前发生退款，此金额也会更新。
		ActualCosPrice      float64 `json:"actualCosPrice"`      //实际计算佣金的金额。订单完成后，会将误扣除的运费券金额更正。如订单完成后发生退款，此金额会更新。
		ActualFee           float64 `json:"actualFee"`           //推客分得的实际佣金（实际计佣金额*佣金比例*最终比例）。如订单完成后发生退款，此金额会更新。
		ValidCode           int     `json:"validCode"`           //订单维度的有效码（-1：未知,2.无效-拆单,3.无效-取消,4.无效-京东帮帮主订单,5.无效-账号异常,6.无效-赠品类目不返佣,7.无效-校园订单,8.无效-企业订单,9.无效-团购订单,10.无效-开增值税专用发票订单,11.无效-乡村推广员下单,12.无效-自己推广自己下单,13.无效-违规订单,14.无效-来源与备案网址不符,15.待付款,16.已付款,17.已完成,18.已结算（5.9号不再支持结算状态回写展示））注：自2018/7/13起，自己推广自己下单已经允许返佣，故12无效码仅针对历史数据有效
		TraceType           int     `json:"traceType"`           //同跨店：2同店 3跨店
		PositionId          int64   `json:"positionId"`          //推广位ID
		SiteId              int64   `json:"siteId"`              //应用id（网站id、appid、社交媒体id）
		UnionAlias          string  `json:"unionAlias"`          //PID所属母账号平台名称（原第三方服务商来源），两方分佣会有该值
		Pid                 string  `json:"pid"`                 //格式:子推客ID_子站长应用ID_子推客推广位ID
		Cid1                int64   `json:"cid1"`                //一级类目id
		Cid2                int64   `json:"cid2"`                //二级类目id
		Cid3                int64   `json:"cid3"`                //三级类目id
		SubUnionId          string  `json:"subUnionId"`          //子渠道标识，在转链时可自定义传入，格式要求：字母、数字或下划线，最多支持80个字符(需要联系运营开放白名单才能拿到数据)
		UnionTag            string  `json:"unionTag"`            //联盟标签数据（32位整型二进制字符串：00000000000000000000000000000001。数据从右向左进行，每一位为1表示符合特征，第1位：红包，第2位：组合推广，第3位：拼购，第5位：有效首次购（0000000000011XXX表示有效首购，最终奖励活动结算金额会结合订单状态判断，以联盟后台对应活动效果数据报表https://union.jd.com/active为准）,第8位：复购订单，第9位：礼金，第10位：联盟礼金，第11位：推客礼金，第12位：京喜APP首购，第13位：京喜首购，第14位：京喜复购，第15位：京喜订单，第16位：京东极速版APP首购，第17位白条首购，第18位校园订单，第19位是0或1时，均代表普通订单，例如：00000000000000000000000000000001:红包订单，00000000000000000000000000000010:组合推广订单，00000000000000000000000000000100:拼购订单，00000000000000000000000000011000:有效首购，00000000000000000000000000000111：红包+组合推广+拼购等） 注：一个订单同时使用礼金和红包，仅礼金位数为1，红包位数为0
		PopId               int64   `json:"popId"`               //商家ID
		Ext1                string  `json:"ext1"`                //推客生成推广链接时传入的扩展字段，订单维度（需要联系运营开放白名单才能拿到数据）
		PayMonth            int     `json:"payMonth"`            //订单维度预估结算时间（格式：yyyyMMdd），0：未结算，订单的预估结算时间仅供参考。账号未通过资质审核或订单发生售后，会影响订单实际结算时间。
		CpActId             int64   `json:"cpActId"`             //招商团活动id：当商品参加了招商团会有该值，为0时表示无活动
		UnionRole           int     `json:"unionRole"`           //站长角色：1 推客 2 团长 3内容服务商
		GiftCouponOcsAmount float64 `json:"giftCouponOcsAmount"` //礼金分摊金额：使用礼金的订单会有该值
		GiftCouponKey       string  `json:"giftCouponKey"`       //礼金批次ID：使用礼金的订单会有该值
		BalanceExt          string  `json:"balanceExt"`          //计佣扩展信息，表示结算月:每月实际佣金变化情况，格式：{20191020:10,20191120:-2}，订单完成后会有该值
		Sign                string  `json:"sign"`                //数据签名，用来核对出参数据是否被修改，入参fields中写入sign时返回
		ProPriceAmount      float64 `json:"proPriceAmount"`      //价保赔付金额：订单申请价保或赔付的金额，实际计佣金额已经减去此金额，您无需处理
		Rid                 int64   `json:"rid"`                 //团长渠道ID，仅限招商团长管理渠道使用，团长开通权限后才可使用。
		GoodsInfo           struct {
			ImageUrl  string `json:"imageUrl"`  //sku主图链接
			Owner     string `json:"owner"`     //g=自营，p=pop
			MainSkuId int64  `json:"mainSkuId"` //自营商品主Id（owner=g取此值
			ProductId int64  `json:"productId"` //非自营商品主Id（owner=p取此值）
			ShopName  string `json:"shopName"`  //店铺名称（或供应商名称）
			ShopId    int64  `json:"shopId"`    //店铺Id
		} `json:"goodsInfo"` //商品信息，入参传入fields，goodsInfo获取
		CategoryInfo struct {
			Cid1     int64  `json:"cid1"`     //一级类目id
			Cid2     int64  `json:"cid2"`     //二级类目id
			Cid3     int64  `json:"cid3"`     //三级类目id
			Cid1Name string `json:"cid1Name"` //一级类目名称
			Cid2Name string `json:"cid2Name"` //二级类目名称
			Cid3Name string `json:"cid3Name"` //三级类目名称
		} `json:"categoryInfo"` //类目信息,入参传入fields，categoryInfo获取
	} `json:"data"`
}

func (J *JdSdk) OrderRowQuery(query interface{}) (res *OrderRowQueryResult) {
	Method := "jd.union.open.order.row.query"
	bodyBytes := J.BodyBytes(Method, query)
	response := &JdUnionOpenOrderRowQueryResponse{}
	e := json.Unmarshal([]byte(bodyBytes), &response)
	if e != nil {
		panic(e)
	}
	res = &OrderRowQueryResult{}
	e = json.Unmarshal([]byte(response.JdUnionOpenOrderRowQueryResponse.Result), res)
	if e != nil {
		panic(e)
	}
	return res
}
