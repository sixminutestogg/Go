package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("chan 可以理解为队列，遵循先进先出的规则。")
	fmt.Println("在说 chan 之前，咱们先说一下 go 关键字。\n\n在 go 关键字后面加一个函数，就可以创建一个线程，函数可以为已经写好的函数，也可以是匿名函数。")

	fmt.Println("Start")

	go func() {
		fmt.Println("Go 语言的线程是并发机制，不是并行机制。")
		fmt.Println("go route meeting")
	}()

	time.Sleep(1 * time.Second)

	fmt.Println("end")

	fmt.Println("声明 chan")
	ch1 := make(chan string)
	fmt.Println("ch1", ch1)

	fmt.Println("声明带10个缓冲的通道")
	ch2 := make(chan string, 10)
	fmt.Println("ch2", ch2)

	fmt.Println("声明只读通道")
	ch3 := make(<-chan string)
	fmt.Println("ch3", ch3)

	fmt.Println("声明只写通道")
	ch4 := make(chan<- string)
	fmt.Println("ch4", ch4)

	fmt.Println("不带缓冲的通道，进和出都会阻塞。\n\n带缓冲的通道，进一次长度 +1，出一次长度 -1，如果长度等于缓冲长度时，再进就会阻塞。")

	fmt.Println("写入 chan")
	ch5 := make(chan string, 10)
	ch5 <- "天地悠悠，过客匆匆，潮起又潮落"
	fmt.Println("ch5", ch5)
	fmt.Println("读取 chan")
	val, ok := <-ch5
	//通道读取一次，第二次读取就没有了
	//val2 := <-ch5
	fmt.Println(val, ok)
	fmt.Println("关闭 chan ")
	close(ch5)

	fmt.Println("使用 close() 通道，安全读取（推荐）")
	ch6 := make(chan string, 10)
	ch6 <- "曹操"
	close(ch6)
	for s := range ch6 {
		fmt.Println(s)
	}

	fmt.Println("配合 goroutine 写入")
	ch7 := make(chan string, 10)

	go func() {
		ch7 <- "我想起了你"
		ch7 <- "过去的时光，不知几何"
		close(ch7)
	}()

	for s := range ch7 {
		fmt.Println(s)
	}

	fmt.Println("- close 以后不能再写入，写入会出现 panic\n- 重复 close 会出现 panic\n- 只读的 chan 不能 close\n- close 以后还可以读取数据")

	fmt.Println("========start=======")
	ch8 := make(chan string)
	go func() {
		fmt.Println("异步读取数据")
		val, ok := <-ch8
		fmt.Println(val, ok)

	}()
	ch8 <- "写入一个数据"
	fmt.Println("==========end==========")

	fmt.Println("========start2=======")
	ch9 := make(chan string, 1)
	ch9 <- "写入一个数据"
	go func() {
		fmt.Println("异步读取数据")
		val, ok := <-ch9
		fmt.Println(val, ok)

	}()
	time.Sleep(1 * time.Second)

	fmt.Println("==========end2==========")

	fmt.Println("------------start-----------")
	ch10 := make(chan string)
	go func() {
		ch10 <- "异步写入一个数据"
	}()
	go func() {
		fmt.Println("异步读取数据")
		val, ok := <-ch10
		fmt.Println(val, ok)
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("----------------end----------")

	ch11 := make(chan string)
	fmt.Println("[producer goroutine] ---> (channel ch) ---> [consumer goroutine]\n")

	go producer(ch11) // 开启生产者
	consumer(ch11)    // 主 goroutine 消费

}
func producer(ch chan<- string) {
	for i := 1; i <= 3; i++ {
		ch <- fmt.Sprintf("任务%d", i)
	}
	close(ch) // 通知对方：我写完了
}

// 消费者函数：只从通道读数据
func consumer(ch <-chan string) {
	for data := range ch {
		fmt.Println("收到:", data)
	}
}
