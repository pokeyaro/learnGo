package main

import "fmt"

func main() {
	/*
		使用 break 关键字，只能一层层的跳
	*/
	// 定义一个标记
	var flagBreak bool
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y == 2 {
				// 设置退出标记
				flagBreak = true
				break
			}
			fmt.Println(x, y)
		}
		// 根据标记, 还需要退出一次循环
		if flagBreak {
			fmt.Println("aa")
			break
		}
	}

	/*
		使用 goto 语句，可任意跳到指定的标签位置
	*/
	for m := 0; m < 10; m++ {
		for n := 0; n < 10; n++ {
			if n == 2 {
				fmt.Println("bb")
				goto label
			}
			fmt.Println(m, n)
		}
	}
label:
	fmt.Println("here")

	/*
		使用 return 语句，虽然可以跳出多层循环，但整个程序都将被终止，无法执行后续代码！🙅
	*/
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				fmt.Println("cc")
				return
			}
			fmt.Println(i, j)
		}
	}
}
