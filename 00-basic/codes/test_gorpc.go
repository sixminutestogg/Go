package main

import (
	"Go/00-basic/codes/friend"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {

	fmt.Println("======== 网络传输 ==========")
	conn, err := grpc.NewClient(
		"192.168.5.158:7000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("无法连接 gRPC: %v", err)
	}
	defer conn.Close()

	fmt.Println("连接成功:", conn.Target())

	fmt.Println("===== 使用老版的连接方式=====")

	fmt.Println("`opts …DialOption`，这个是不定参数传递，参数的类型为 `DialOption`，不定参数是指函数传入的参数个数为不定数量，可以不传，也可以为多个。")

	fmt.Println("写一个不定参数传递的方法也很简单，看看下面这个方法 1 + 2 + 3 = 6。")

	fmt.Println(Add(1, 2, 3))

	fmt.Println("`WithInsecure()`、`WithBlock()` 类似于这样的 With 方法，其实作用就是修改 `dialOptions` 结构体的配置，之所以这样写我个人认为是面向对象的思想，当配置项调整的时候调用方无需修改。")

	fmt.Println("模拟一个场景，使用 `不定参数` 和 `WithXXX` 这样的写法，写个 Demo，比如我们要做一个从附近找朋友的功能，" +
		"配置项有：性别、年龄、身高、体重、爱好，我们要找性别为女性，年龄为30岁，身高为160cm，体重为55kg，爱好为爬山的人，希望是这样的调用方式：")

	find, err := friend.Find("今天干什么",
		friend.WithSex(1),
		friend.WithAge(30),
		friend.WithHeight(18),
		friend.WithHobby("爬山"),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(find)
}

func Add(a int, args ...int) (result int) {
	result += a
	for _, arg := range args {
		result += arg
	}
	return
}
