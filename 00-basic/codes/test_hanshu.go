package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"time"
)

type ApiResult struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func main() {
	fmt.Println("函数")

	var res ApiResult

	res.Code = 200
	res.Message = "操作成功"
	res.Data = map[string]interface{}{
		"天": "订单",
		"地": "上述",
	}
	fmt.Println("res:", res)

	toJsonApi(&res)

	setApiData(&res)

	toJsonApi(&res)

	str := "123245"

	fmt.Printf("MD5(%s):%s\n", str, Md5(str))

	fmt.Printf("current time str : %s\n", getTimeStrApi())

	fmt.Printf("current time str : %s\n", getTimeStrApi())
	fmt.Printf("current time int : %d\n", getTimeIntApi())
	mapApi := map[string]interface{}{
		"one":   "第一天",
		"two":   "第二天",
		"three": "第三天",
	}
	fmt.Printf("自定义密码签名：%s\n", creatSignApi(mapApi))
}

func setApiData(res *ApiResult) {
	res.Code = 500
	res.Message = "操作错误"
}

func toJsonApi(res *ApiResult) {
	marshal, err := json.Marshal(res)
	if err != nil {
		fmt.Println(" error = ", err)
	}
	fmt.Println("json format", string(marshal))

}

func Md5(password string) string {
	//md5.New() 返回一个实现了 hash.Hash 接口的对象。
	s := md5.New()
	//Write() 方法向哈希计算器输入数据。
	//[]byte(password) 是把字符串转换为字节数组（因为哈希算法处理的是二进制数据）。
	s.Write([]byte(password))
	//Sum(nil) 会计算出当前输入数据的 MD5 摘要结果
	//由于 MD5 的原始输出是二进制的，我们通常把它转成十六进制字符串显示。
	//hex.EncodeToString() 就是把字节数组转成常见的 32 位十六进制字符串。
	return hex.EncodeToString(s.Sum(nil))
}

// 获取当前时间字符串
func getTimeStrApi() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 获取当前时间戳
func getTimeIntApi() int64 {
	return time.Now().Unix()
}

// 生成签名方法
func creatSignApi(params map[string]interface{}) string {
	var key []string
	var str = ""
	for k := range params {
		key = append(key, k)
	}
	sort.Strings(key)

	for i := 0; i < len(key); i++ {
		if i == 0 {
			str = fmt.Sprintf("%v=%v", key[i], params[key[i]])
		} else {
			str = str + fmt.Sprintf("&xl_%v=%v", key[i], params[key[i]])
		}

	}

	// 自定义密钥
	var secretKey = "天地悠悠，过客匆匆"
	// 自定义签名算法
	sign := Md5(Md5(str) + Md5(secretKey))

	return sign
}
