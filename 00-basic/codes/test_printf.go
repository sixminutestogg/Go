package main

import (
	"fmt"
	"strings"
	"testing"
)

func main() {
	fmt.Println("不要使用 `+` 和 `fmt.Sprintf` 操作字符串，虽然很方便，但是真的很慢！\n\n我们要使用 `bytes.NewBufferString` 进行处理。")

}

func BenchmarkStringOperation1(b *testing.B) {
	b.ResetTimer()
	str := ""
	for i := 0; i < b.N; i++ {
		str += "golang"
	}

}

func BenchmarkStringOperation2(b *testing.B) {
	b.ResetTimer()
	var builder strings.Builder
	for i := 0; i < b.N; i++ {
		builder.WriteString("golang")
	}
}
