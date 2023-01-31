package main

import "fmt"

func main() {
	// 使用转义符
	msg1 := "hello ~\r\n\"naruto\""

	// 使用多行字符串
	msg2 := `hello ~
"naruto"`

	fmt.Println(msg1)
	fmt.Println(msg2)
}
