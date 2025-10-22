package main

import "fmt"

func main() {

	person := [3]string{"Tom", "Jack", "Jerry"}
	fmt.Printf("len=%d cap=%d array=%v\n", len(person), cap(person), person)

	fmt.Println("循环")

	for k, v := range person {
		fmt.Printf("person[%d]:%s\n", k, v)
	}

	fmt.Println("空白格，表示不关心此值")
	for _, k := range person {
		fmt.Println(k)
	}

	for i := range person {
		fmt.Printf("person[%d]: %s\n", i, person[i])
	}
	fmt.Println("======================")

	for i := 0; i < len(person); i++ {
		fmt.Printf("person[%d]:%s\n", i, person[i])
	}
	fmt.Println("------------------------")

	fmt.Println("循环 slice")

	fmt.Println("在语法上 for range 遍历数组和切片几乎一样，\n但数组遍历时会复制整个数组，而切片不会。\n因此——实际开发中几乎总是使用切片而不是数组。")

	personMap := map[string]interface{}{
		"人": "tom",
		"日": "太乙",
		"月": "月光",
	}
	fmt.Println("personMap", personMap)

	for key, value := range personMap {
		fmt.Printf("personMap[%d]:%s\n", key, value)
	}

	fmt.Println("break")
	for i := 0; i < 10; i++ {
		if i == 6 {
			break
		}
		fmt.Println("i=", i)
	}

	fmt.Println("continue")
	for i := 0; i < 10; i++ {
		if i == 6 {
			continue
		}
		fmt.Println("i=", i)
	}

	fmt.Println("goto 改变函数内代码执行顺序，不能跨函数使用。")

	for i := 0; i < 10; i++ {
		if i == 6 {
			goto END
		}
		fmt.Println("i =", i)
	}
END:
	fmt.Println("end")

	fmt.Println("switch")

	i := 4
	fmt.Printf("当i = %d", i)

	switch i {
	case 1:
		fmt.Println("输出i = ", i)
	case 2:
		fmt.Println("输出 i =", 2)
	case 3:
		fmt.Println("输出 i =", 3)
		fallthrough
	case 4, 5, 6:
		fmt.Println("输出 i =", "4 or 5 or 6")
	default:
		fmt.Println("输出 i =", "xxx")
	}

	fmt.Println("- 默认每个 case 带有 break\n- case 中可以有多个选项\n- fallthrough 不跳出，并执行下一个 case")
}
