package main

import "fmt"

func a() (int, bool) {
	return 0, true
}

func main() {
	var _ int    // 定义匿名变量
	_, ok := a() // 接收匿名变量
	if ok {
		// 打印
		fmt.Println("ok")
	}
}
