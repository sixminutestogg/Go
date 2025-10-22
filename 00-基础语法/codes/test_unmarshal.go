package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {

	jsonStr := `{"number":1234567}`
	fmt.Println("json:", jsonStr)

	result := make(map[string]interface{})

	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		panic(err)
	}
	fmt.Println("当数字的位数大于 6 位时，才会变成了科学计数法")
	fmt.Println("result:", result)
	fmt.Println("解决方法一：强制类型转换")
	fmt.Println(int(result["number"].(float64)))
	fmt.Println("解决方法二：尽量避免使用 `interface`，对 `json` 字符串结构定义结构体,使用在线地址 ：https://mholt.github.io/json-to-go/")

	type Num struct {
		Number int `json:"number"`
	}
	var num Num
	err2 := json.Unmarshal([]byte(jsonStr), &num)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(num)

	fmt.Println("解决方法三：使用 `UseNumber()` 方法。")
	decoder := json.NewDecoder(bytes.NewReader([]byte(jsonStr)))
	decoder.UseNumber()
	err = decoder.Decode(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	marshal, _ := json.Marshal(result)
	fmt.Println(string(marshal))

}
