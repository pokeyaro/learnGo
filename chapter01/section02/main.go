package main

import "fmt"

func main() {
	// 显式定义
	const PI float32 = 3.1415926

	// 定义一组
	const (
		UNKNOWN = 100
		SUCCESS = 200
		FAILED  = 500
	)

	// 省略的方式，如果后续不定义类型和值，就会延续之前定义的类型和值，与之前保持一致
	const (
		P1 = "鸣人"
		P2
		P3
		U1 = "佐助"
		U2
	)
	fmt.Println(P1, P2, P3, U1, U2)
}
