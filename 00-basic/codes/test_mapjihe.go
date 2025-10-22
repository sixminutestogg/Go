package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Map 集合是无序的 key-value 数据结构。")
	fmt.Println("Map 集合中的 key / value 可以是任意类型，" +
		"但所有的 key 必须属于同一数据类型，" +
		"所有的 value 必须属于同一数据类型，key 和 value 的数据类型可以不相同。")

	// 声明一个 Map
	var p1 map[int]string
	p1 = make(map[int]string)
	p1[1] = "Tom"
	p1[2] = "Jerry"
	p1[3] = "Alice"
	p1[4] = "Team"
	fmt.Println("p1 = ", p1)

	var p2 map[int]string = map[int]string{}
	p2[1] = "Tom"
	p2[2] = "Jerry"
	p2[3] = "Tian"
	fmt.Println("p2 = ", p2)

	var p3 map[int]string = make(map[int]string)
	p3[1] = "Tom"
	p3[2] = "Jerry"
	p3[3] = "Tao"
	fmt.Println("p3 = ", p3)

	p4 := map[int]string{}
	p4[1] = "Tom"
	fmt.Println("p4=", p4)

	p5 := make(map[int]string)
	p5[1] = "Tom"
	p5[2] = "曹操"
	fmt.Println("p5=", p5)

	p6 := map[int]string{
		1: "Tom",
		2: "Jerty",
	}
	fmt.Println("p6=", p6)

	//生成 JSON
	fmt.Println("生成 JSON")

	res := make(map[string]interface{})
	res["code"] = 200
	res["msg"] = "success"
	res["data"] = map[string]interface{}{
		"username": "Tom",
		"age":      "30",
		"hobby":    []string{"读书", "爬山"},
	}
	fmt.Println("map data : ", res)

	//序列化
	marshal, err := json.Marshal(res)
	if err != nil {
		fmt.Println("json error", err)
	}
	fmt.Println("json string :", marshal)
	fmt.Println("")
	fmt.Println("map to json")
	fmt.Println("json data ", string(marshal))

	//反序列化
	res2 := make(map[string]interface{})
	errfan := json.Unmarshal(marshal, &res2)
	if errfan != nil {
		fmt.Println("json marshal error:", errfan)
	}
	fmt.Println("")
	fmt.Println("--- json to map ---")
	fmt.Println("map data :", res2)
	fmt.Println("map data :", errfan)

	fmt.Println("编辑和删除")

	person := map[int]string{
		1: "天",
		2: "地",
		3: "人",
		4: "合",
	}
	fmt.Println("data", person)

	delete(person, 2)
	fmt.Println("data", person)

	person[3] = "jack"
	person[4] = "Tome"
	person[2] = "丹"
	fmt.Println("data", person)

}
