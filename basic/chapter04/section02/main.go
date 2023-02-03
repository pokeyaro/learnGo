package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	 * for init; condition; post {}
	 */

	// 1-100 相加
	var sum int
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("1 + ... + 100 =", sum)

	// 经典99乘法表
	/*
			1*1=1
			1*2=2 2*2=4
			1*3=3 2*3=6 3*3=9
		    ...
	*/
	for row := 1; row <= 9; row++ {
		for col := 1; col <= row; col++ {
			fmt.Printf("%dx%d=%d\t", col, row, col*row)
		}
		fmt.Println()
	}

	// break 与 continue
	for round := 0; ; round++ {
		time.Sleep(1 * time.Second)
		if round > 10 {
			fmt.Println("exit")
			break
		}
		if round%2 == 0 {
			fmt.Println("skip")
			continue
		}
		fmt.Println(round)
	}
}
