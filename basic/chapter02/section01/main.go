package main

import "fmt"

func main() {
	var b bool = true

	var n int = 10

	var f1 float32 = 3.14
	var f2 float64 = 3.1415

	var c1 byte = 'a' + 1 // 适合存放ascii字符，可以进行运算，本质上就是整数
	var c2 rune = '啊'     // 存放utf-8字符

	var s string = "abc"

	var c int32 = 128522
	fmt.Println(b, n, f1, f2, c1, c2, s)
	fmt.Printf("%v, %c", c, c)
}
