package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type PersonTest struct {
	Name    string
	Age     int
	Address string
	Sex     bool
}

type Order struct {
	Name     string
	Price    int
	Discount int
}

func main() {

	fmt.Println("声明结构体")

	var p1 PersonTest

	p1.Name = "张三"
	p1.Age = 30
	p1.Address = "北京"
	p1.Sex = true
	fmt.Println("p1=", p1)

	var p2 = PersonTest{Name: "李四", Age: 40, Address: "上海", Sex: false}

	fmt.Println("p2=", p2)

	p3 := PersonTest{Name: "王五", Age: 50, Address: "广州", Sex: true}

	fmt.Println("p3=", p3)

	//匿名结构体
	p4 := struct {
		Name string
		Sex  bool
	}{Name: "d", Sex: true}

	fmt.Println("p4=", p4)

	//生成 JSON
	type Result struct {
		Code    int    `json:"code"`
		Message string `json:"msg"`
	}

	var res Result
	res.Code = 200
	res.Message = "success"

	//序列化
	jsons, errs := json.Marshal(res)
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	fmt.Println("json data :", string(jsons))

	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}

	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	marshal, errs := json.Marshal(group)
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	fmt.Println("json data :", string(marshal))
	os.Stdout.Write(marshal)

	fmt.Println("----------------------------------")

	//反序列化
	var res2 Result
	errorTian := json.Unmarshal(jsons, &res2)
	if errorTian != nil {
		fmt.Println("json unmarshal error:", errorTian)
	}
	fmt.Println("res2 :", res2)

	var jsonBlob = []byte(`[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll",    "Order": "Dasyuromorphia"}
	]`)

	type Animal struct {
		Name  string
		Order string
	}

	var animals []Animal

	erranimals := json.Unmarshal(jsonBlob, &animals)
	if erranimals != nil {
		fmt.Println("json unmarshal error:", erranimals)
	}
	fmt.Println("err : ", erranimals)
	fmt.Println("json : ", animals)

	//改变数据
	fmt.Println("改变数据")

	order := Order{
		Name:     "洗浴",
		Price:    9000,
		Discount: 1000,
	}

	toJsonOrder(&order)

	setOrderData(&order)
	toJsonOrder(&order)

}

func setOrderData(order *Order) {
	order.Discount = 2000
	order.Price = 1000
	order.Name = "按摩"
}

func toJsonOrder(order *Order) {
	marshal, err := json.Marshal(order)
	if err != nil {
		fmt.Println("json marshal error:", err)
	}
	fmt.Println("json data :", string(marshal))
}
