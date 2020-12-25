package JdunionSdk

import (
	"encoding/json"
)

type OrderBonusQueryResult struct {
	Code      int    `json:"code"`
	HasMore   bool   `json:"hasMore"`
	Message   string `json:"message"`
	RequestId string `json:"requestId"`
	Data      []struct {
		UnionId          int64   `json:"unionId"`          //联盟ID
		BonusInvalidCode string  `json:"bonusInvalidCode"` //无效状态码，-1:无效、2:无效-拆单、3:无效-取消、4:无效-京东帮帮主订单、5:无效-账号异常、6:无效-赠品类目不返佣 等
		BonusInvalidText string  `json:"bonusInvalidText"` //无效状态码对应的无效状态文案
		PayPrice         float64 `json:"payPrice"`         //实际支付金额
		EstimateCosPrice float64 `json:"estimateCosPrice"` //预估计佣金额
		EstimateFee      float64 `json:"estimateFee"`      //预估佣金
		ActualCosPrice   float64 `json:"actualCosPrice"`   //实际计佣金额
		ActualFee        float64 `json:"actualFee"`        //实际佣金
		OrderTime        int64   `json:"orderTime"`        //下单时间、时间戳（毫秒）
		FinishTime       int64   `json:"finishTime"`       //完成时间、时间戳（毫秒）
		PositionId       int64   `json:"positionId"`       //推广位ID
		OrderId          int64   `json:"orderId"`          //订单号
		BonusState       int     `json:"bonusState"`       //奖励状态，0:无效、1:有效
		BonusText        string  `json:"bonusText"`        //奖励状态文案
		SkuName          string  `json:"skuName"`          //商品名称
		CommissionRate   float64 `json:"commissionRate"`   //佣金比例，单位：%
		SubUnionId       string  `json:"subUnionId"`       //子联盟ID
		Pid              string  `json:"pid"`              //pid
		Ext1             string  `json:"ext1"`             //推客生成推广链接时传入的扩展字段
		UnionAlias       string  `json:"unionAlias"`       //母站长简称
		SubSideRate      float64 `json:"subSideRate"`      //分成比例（单位：%）
		SubsidyRate      float64 `json:"subsidyRate"`      //补贴比例（单位：%）
		FinalRate        float64 `json:"finalRate"`        //最终分佣比例（单位：%）=分成比例+补贴比例
		ActivityName     string  `json:"activityName"`     //活动名称
		ParentId         int64   `json:"parentId"`         //父单的订单号：如一个订单拆成多个子订单时，原订单号会作为父单号，拆分的订单号为子单号存储在orderid中。若未发生拆单，该字段为0
		SkuId            int64   `json:"skuId"`            //skuId
		EstimateBonusFee float64 `json:"estimateBonusFee"` //预估奖励金额：查询时间范围内，已付款且奖励有效，满足奖励规则的奖励金额
		ActualBonusFee   float64 `json:"actualBonusFee"`   //实际奖励金额：查询时间范围内，已付款或已完成（视具体规则），奖励有效且满足奖励规则的奖励金额
		OrderState       int     `json:"orderState"`       //奖励订单状态，1:已完成、2:已付款、3:待付款
		OrderText        string  `json:"orderText"`        //奖励订单状态，待付款/已付款/已完成
		SortValue        string  `json:"sortValue"`        //排序值，按'下单时间'分页查询时使用
	} `json:"data"`
}

type JdUnionOpenOrderBonusQueryResponse struct {
	JdUnionOpenOrderBonusQueryResponse struct {
		Result string `json:"result"`
		Code   string `json:"code"`
	} `json:"jd_union_open_order_bonus_query_response"`
}

type OrderBonusQueryRequest struct {
	OrderReq struct {
		OptType   int    `json:"optType"`   //时间类型（1.下单时间拉取、2.更新时间拉取）
		StartTime int64  `json:"startTime"` //订单开始时间，时间戳（毫秒），与endTime间隔不超过10分钟
		EndTime   int64  `json:"endTime"`   //订单结束时间，时间戳（毫秒），与startTime间隔不超过10分钟
		PageNo    int    `json:"pageNo"`    //页码，默认值为1
		PageSize  int    `json:"pageSize"`  //每页数量，上限100
		SortValue string `json:"sortValue"` //与pageNo、pageSize组合使用。获取当前页最后一条记录的sortValue，下一页请求传入该值
	} `json:"orderReq"`
}

func (J *JdSdk) OrderBonusQuery(query interface{}) (result *OrderBonusQueryResult) {
	Method := "jd.union.open.order.bonus.query"
	bodyBytes := J.BodyBytes(Method, query)
	response := &JdUnionOpenOrderBonusQueryResponse{}
	e := json.Unmarshal([]byte(bodyBytes), &response)
	if e != nil {
		panic(e)
	}
	result = &OrderBonusQueryResult{}
	e = json.Unmarshal([]byte(response.JdUnionOpenOrderBonusQueryResponse.Result), result)
	if e != nil {
		panic(e)
	}
	return result
}
