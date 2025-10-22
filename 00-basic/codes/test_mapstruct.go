package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Family struct {
	LastName string
}

type Location struct {
	City string
}

type PersonMap struct {
	Family    `mapstructure:",squash"`
	Location  `mapstructure:",squash"`
	FirstName string
}

func main() {
	intput := map[string]interface{}{
		"FirstName": "天",
		"LastName":  "史密斯",
		"City":      "新西兰",
	}

	fmt.Println("input :", intput)

	var result PersonMap

	err := mapstructure.Decode(intput, &result)
	if err != nil {
		panic(err)
	}
	fmt.Println("使用的是 mapstructure 包，struct tag 标识不要写 json，要写 mapstructure。")

	fmt.Println(result.FirstName)
	fmt.Println(result.LastName)
	fmt.Println(result.City)

}
