package JdunionSdk

import (
	"encoding/json"
	"strings"
)

/*商品类型返回数据结构体*/
//start
type CategoryResponse struct {
	JdUnionOpenCategoryGoodsGetResponse struct {
		Result string `json:"result"`
		Code   string `json:"code"`
	} `json:"jd_union_open_category_goods_get_response"`
}

type CateGoryResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		Grade    int    `json:"grade"`
		ParentId int    `json:"parentId"`
	} `json:"data"`
	RequestId string `json:"requestId"`
}

//end

//获取京东商品类型列表

func (J *JdSdk) GetCategoryList(UriQuery string) (categoryResult *CateGoryResult) {
	Method := "jd.union.open.category.goods.get"
	J.SetSignJointUrlParam(Method, UriQuery)
	var urls strings.Builder
	urls.WriteString(J.JdHost)
	urls.WriteString(J.SignAndUri)
	body, _ := HttpGet(urls.String())
	result := &CategoryResponse{}
	e := json.Unmarshal([]byte(body), &result)
	if e != nil {
		panic(e)
	}
	categoryResult = &CateGoryResult{}
	e = json.Unmarshal([]byte(result.JdUnionOpenCategoryGoodsGetResponse.Result), categoryResult)
	if e != nil {
		panic(e)
	}
	return categoryResult
}

//请求业务参数

type Req struct {
	ParentId int64 `json:"parentId"` //父类目id(一级父类目为0)
	Grade    int64 `json:"grade"`    //类目级别(类目级别 0，1，2 代表一、二、三级类目)
}
