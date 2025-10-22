package main

import (
	"fmt"
)

func main() {

	/**
	 * 常量定义
	 */

	const name string = "Tom"
	fmt.Println(name)

	const age = 30
	fmt.Println(age)

	const name_1, name_2 string = "Tom", "Jay"
	fmt.Println(name_1, name_2)

	const name_3, age_1 = "Tom", 30
	fmt.Println(name_3, age_1)

	const remark, total string = "dd", "督导"

	fmt.Println(remark, total)

	const bootAge = 19
	fmt.Println(bootAge)

	const one, two, three = "天", "地", "人"

	fmt.Println(one, two, three)

}
