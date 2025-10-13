package main

import (
	"fmt"
)

func main() {

	//常量声明
	const name string = "tom"

	fmt.Println(name)

	const age = 30

	fmt.Println(age)

	const name_1, name_2, name_3 string = "tom", "jack", "mary"

	fmt.Println(name_1, name_2, name_3)

	const name_end, age_time = "tom", 30

	fmt.Println(name_end, age_time)

	//变量声明
	//**单个变量声明**
	//第一种：var 变量名称 数据类型 = 变量值
	var go_name string = "tian"
	fmt.Println(go_name)
	//第二种：var 变量名称 = 变量值
	var go_age = 30
	fmt.Println(go_age)

	//第三种：变量名称 := 变量值
	marray_name := "商博涵"
	fmt.Println(marray_name)

	//**多个变量声明**
	//第一种：var 变量名称,变量名称 ... ,数据类型 = 变量值,变量值 ...
	var one, two, three, four string = "one", "two", "three", "four"
	fmt.Println(one, two, three, four)

	//第二种：var 变量名称,变量名称 ...  = 变量值,变量值 ...
	var tain, di, ren, he = "天", "地", "人", "和"
	fmt.Println(tain, di, ren, he)

	//第三种：变量名称,变量名称 ... := 变量值,变量值 ...
	zi, chou, ying, mou := "子", "丑", "寅", "卯"
	fmt.Println(zi, chou, ying, mou)

	//输出方法

	//**fmt.Print**：输出到控制台（仅只是输出）
	fmt.Print("输出到控制台不换行")
	fmt.Println("----")
	//**fmt.Println**：输出到控制台并换行
	fmt.Println("输出到控制台并换行")
	fmt.Println("----")

	//**fmt.Printf**：仅输出格式化的字符串和字符串变量（整型和整型变量不可以）
	fmt.Printf("输出格式化的字符串和字符串变量：%s\n", "商博涵")
	fmt.Printf("name=%s,age=%d\n", "商博涵", 30)
	//**fmt.Sprintf**：格式化并返回一个字符串，不输出。
	fmt.Printf("name=%s,age=%d,height=%v\n", "商博涵", 30, fmt.Sprintf("%.2f", 180.567))
	fmt.Printf("输出格式化的字符串和字符串变量：%s\n", fmt.Sprintf("商博涵"))

}
