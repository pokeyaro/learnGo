package main

import "fmt"

func main() {
	// 运算符
	var a, b = 5, 3
	fmt.Println(a + b) // 加 8
	fmt.Println(a - b) // 减 2
	fmt.Println(a * b) // 乘 15
	fmt.Println(a / b) // 除 1
	fmt.Println(a % b) // 余 2

	// + 号：字符串拼接
	var s1, s2 = "Hello", "World"
	s := fmt.Sprintf("%v", s1+s2)
	fmt.Println(s) // HelloWorld

	// 自增或自减
	var i int // 0
	i++       // 1
	i += 1    // 2
	i = i + 1 // 3
	i = i - 1 // 2
	i -= 1    // 1
	i--       // 0

	// 逻辑运算符
	var v1, v2 = true, false
	if v1 && v2 {
		fmt.Println("always false")
	}
	if v1 || v2 {
		fmt.Println("always true")
	}
	if !v1 {
		fmt.Println("always false")
	}

	// 位运算符（一般对性能要求高的时候，才会考虑位运算）
	var n1 = 60          // 0011 1100
	var n2 = 13          // 0000 1101
	fmt.Println(n1 & n2) // 0000 1100 -> 与运算   12 (2^3 + 2^2)
	fmt.Println(n1 | n2) // 0011 1101 -> 或运算   61 (2^5 + 2^4 + 2^3 + 2^2 + 2^0)
	fmt.Println(n1 ^ n2) // 0011 0001 -> 异或运算 49 (2^5 + 2^4 + 2^0)
}
