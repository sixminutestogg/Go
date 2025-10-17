package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type MobileInfo struct {
	Resultcode string `json:"resultcode"`
	Reason     string `json:"reason"`
	Result     struct {
		Province string `json:"province"`
		City     string `json:"city"`
		Areacode string `json:"areacode"`
		Zip      string `json:"zip"`
		Company  string `json:"company"`
		Card     string `json:"card"`
	} `json:"result"`
}

func main() {
	jsonStr := `{
    "resultcode": "200",
    "reason": "Return Successd!",
    "result": {
        "province": "浙江",
        "city": "杭州",
        "areacode": "0571",
        "zip": "310000",
        "company": "中国移动",
        "card": ""
   		 }
	}`

	var mobile MobileInfo
	err := json.Unmarshal([]byte(jsonStr), &mobile)
	if err != nil {
		fmt.Print("error", err.Error())
	}
	fmt.Println(mobile.Resultcode)
	fmt.Println(mobile.Reason)
	fmt.Println(mobile.Result.City)

	fmt.Printf("mobile json is %+v\n:", mobile)

	marshal2, _ := json.Marshal(mobile)
	fmt.Println("紧凑 json : ", string(marshal2))

	indent, _ := json.MarshalIndent(mobile, "", " ")
	fmt.Println("格式化 json : \n", string(indent))

	fmt.Println("============= 方法体中 ，定义的类型是字符串，接到的确实int  =========================")

	jsonStrOne := `{
			"resultcode": 200
		}`

	//1.先用map接所有
	var mapResult map[string]interface{}
	errMap := json.Unmarshal([]byte(jsonStrOne), &mapResult)
	if errMap != nil {
		fmt.Println("error is result : ", errMap.Error())
	}
	fmt.Println("mapResult", mapResult)

	//2.弱类型解析的方法 `WeakDecode()` 转换

	var mobileOne MobileOne
	err = mapstructure.WeakDecode(mapResult, &mobileOne)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(mobileOne.ResultCode)

	fmt.Println(mobileOne)

}

type MobileOne struct {
	ResultCode string `json:"resultCode"`
}
