package Jdunion

import (
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"reflect"
	"sort"
	"strings"
	"time"
)

//生成请求参数和签名
func (J *JdSdk) SetSignJointUrlParam(Method string, ParamJson string) *JdSdk {
	J.RequestParam.Method = Method
	location, _ := time.LoadLocation("Asia/Shanghai")
	J.RequestParam.Timestamp = time.Now().In(location).Format("2006-01-02 15:04:05")
	J.RequestParam.Param_json = ParamJson
	values := reflect.ValueOf(J.RequestParam)
	keys := reflect.TypeOf(J.RequestParam)
	count := values.NumField()
	SortSlice := Persons{}
	for i := 0; i < count; i++ {
		value := values.Field(i)
		key := keys.Field(i)
		if key.Name == "Param_json" {
			key.Name = "360buy_param_json"
		}
		switch value.Kind() {
		case reflect.String:
			//if key.Name == "Param_json"{
			//	key.Name = "360buy_param_json"
			//}
			SortSlice = append(SortSlice, Onestruct{strings.ToLower(key.Name), value.String()})
		case reflect.Int:
			SortSlice = append(SortSlice, Onestruct{strings.ToLower(key.Name), value.String()})
		}
	}
	//fmt.Println(SortSlice)
	sort.Sort(SortSlice)
	var builder strings.Builder
	u := url.Values{}
	builder.WriteString(J.AppSecret)
	for _, person := range SortSlice {
		if person.Value == "" {
			continue
		}
		builder.WriteString(strings.ToLower(person.Key) + person.Value)
		u.Add(strings.ToLower(person.Key), person.Value)
	}
	builder.WriteString(J.AppSecret)
	//生成签名
	u.Add("sign", strings.ToUpper(Md5(builder.String())))
	//拼接参数
	J.SignAndUri = u.Encode()
	return J
}

//生成 MD5
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
