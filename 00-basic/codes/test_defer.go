package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	fmt.Println("延迟执行（defer execution）\n让某个函数（或表达式）在 当前函数返回之前 自动执行。")
	defer someFunction()
	fmt.Println("不立即执行 someFunction()，\n而是在 外层函数（当前函数）结束时 执行它。")

	fmt.Println("A")
	defer fmt.Println("B")
	fmt.Println("c")

	defer fmt.Println("第一天")
	defer fmt.Println("第二天")
	defer fmt.Println("第三天")

	fmt.Println("多个 defer：先进后出 (LIFO)")

	fmt.Println("典型使用场景")
	fmt.Println("文件关闭")
	dir, _ := os.Getwd()
	fmt.Println("当前工作目录：", dir)

	open, err := os.Open("README.md")
	if err != nil {
		log.Fatalln(err)
	}
	name := open.Name()
	fmt.Println(name)
	defer open.Close()

	x := 1
	y := 2
	defer calc("A", x, calc("B", x, y))
	x = 3
	defer calc("C", x, calc("D", x, y))
	y = 4

	var a = 1
	var b = 2
	defer fmt.Println("-------main --------", a+b)
	defer func() {
		fmt.Println("-------main 2--------:", a+b)
	}()
	defer func(a int, b int) {
		fmt.Println("----------main 3 ------ :", a+b)
	}(a, b)
	a = 3
	fmt.Println("main")

	fmt.Println("==========  return 操作 ===============")

	GoA()
	time.Sleep(1 * time.Second)
	fmt.Println("main two ")
	fmt.Println("同一个 goroutine 里的 panic\t✅\tdefer 和 panic 在同一调用栈\n不同 goroutine 的 panic\t❌\trecover() 只能作用于当前 goroutine")
}

func someFunction() {
	fmt.Println("一些方法")
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b)
	return ret
}

func GoA() {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("panic: " + fmt.Sprintf("%s", err))
			}
		}()
		GoB()
	}()

}

func GoB() {
	panic("error")
}
