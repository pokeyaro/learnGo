package main

import "fmt"

func main() {
	const (
		ERR01 = iota
		ERR02
		ERR03
	)
	fmt.Println(ERR01, ERR02, ERR03) // 0 1 2

	const (
		D1 = iota + 1 // 1
		D2            // 2
		D3            // 3
		S1 = "abc"    // abc
		S2            // abc
		S3            // abc
		D4 = iota     // 6
		D5            // 7
	)
	fmt.Println(D1, D2, D3, D4, D5, S1, S2, S3)
}
