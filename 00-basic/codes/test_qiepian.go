package main

import (
	"fmt"
)

func main() {
	fmt.Println("切片测试")

	//声明切片
	//nil 切片
	var sli1 []int
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli1), cap(sli1), sli1)

	//空切片
	var sli2 = []int{}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli2), cap(sli2), sli2)

	//数组
	var sli3 = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli3), cap(sli3), sli3)

	sil4 := [5]int{4, 5, 6, 1, 3}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sil4), cap(sil4), sil4)

	//t Type: 要创建的类型（可以是切片、映射或通道）
	//size ...IntegerType: 可变参数，表示类型的大小/长度
	//对于切片：make([]int, length, capacity)
	//对于映射：make(map[keyType]valueType)
	//对于通道：make(chan type, buffer)
	var sli5 []int = make([]int, 5, 8)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli5), cap(sli5), sli5)

	// 切片示例
	s := make([]int, 5, 10) // 长度为5，容量为10的切片
	fmt.Println("切片:", s)   // 输出: [0 0 0 0 0]

	// 映射示例
	m := make(map[string]int) // 字符串到整数的空映射
	m["apple"] = 5
	fmt.Println("映射:", m) // 输出: map[apple:5]

	// 通道示例
	c := make(chan int, 5) // 缓冲容量为5的通道
	fmt.Println("通道:", c)  // 输出: 0x...（通道的内存地址）

	// nil值示例
	var ns []int
	fmt.Println("nil切片:", ns)         // 输出: []
	fmt.Println("是否为nil:", ns == nil) // 输出: true

	sli6 := make([]int, 5, 10)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli6), cap(sli6), sli6)

	//截取切片

	sli7 := []int{78, 3, 32, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli7), cap(sli7), sli7)

	fmt.Println("sli7[1]=", sli7[1])
	fmt.Println("sli7[:]=", sli7[:])
	fmt.Println("sli7[1:3]=", sli7[1:3])
	fmt.Println("sli7[1:]=", sli7[1:])
	fmt.Println("sli7[:4]=", sli7[:4])
	fmt.Println("sli7[:1]=", sli7[:1])
	fmt.Println("sli7[:10]=", sli7[:10])

	//追加切片
	sli8 := []int{7, 3, 2, 45, 3}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli8), cap(sli8), sli8)

	sli8 = append(sli8, 5)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli8), cap(sli8), sli8)

	sli8 = append(sli8, 6, 7, 8)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli8), cap(sli8), sli8)

	//删除切片

	sli9 := []int{7, 3, 2, 45, 3, 8, 98}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli9), cap(sli9), sli9)

	//删除尾部 2 个元素
	sli9 = sli9[:len(sli9)-2]
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli9), cap(sli9), sli9)

	//删除开头 2 个元素
	sli9 = sli9[2:]
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli9), cap(sli9), sli9)

	//删除中间 2 个元素
	sli9 = append(sli9[:3])
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli9), cap(sli9), sli9)

}
