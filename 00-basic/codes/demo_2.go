package main

import (
	"fmt"
)

/**
 * @Description: 变量声明
 */
func main() {

	var age_1 uint8 = 31
	var age_2 = 32
	age_3 := 33
	fmt.Println(age_1, age_2, age_3)

	var age_4, age_5, age_6 int = 31, 32, 33
	fmt.Println(age_4, age_5, age_6)

	var name_1, age_7 = "Tom", 30
	fmt.Println(name_1, age_7)

	name_2, is_boy, height := "Jay", true, 180.66
	fmt.Println(name_2, is_boy, height)

	var blName string = "添加"
	fmt.Println(blName)

	var tom = "tom simith"
	fmt.Println(tom)

	difjd := 9
	fmt.Println(difjd)

	var one, two, three int = 1, 2, 3
	fmt.Println(one, two, three)

	var name_3, name_5, name_6 = "Tom", "Jay", "Jack"

	fmt.Println(name_3, name_5, name_6)

	day, month, year, name := 1, 2, 3, "Tom"

	fmt.Println(day, month, year, name)

}
