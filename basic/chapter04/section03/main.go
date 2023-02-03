package main

import "fmt"

func main() {
	// utf-8 中英文混合字符串
	var cartoon string = "Naruto 火影忍者 🍃"

	// 这种方式是错误的 🙅‍
	for k, _ := range cartoon {
		fmt.Printf("%c\n", cartoon[k]) // 一个中文占3个字节，这样只会取出第1个字节，中文会出现乱码！
	}

	// 方式1 ✅
	for _, v := range cartoon {
		fmt.Printf("%c\r\n", v)
	}

	// 方式2 ✅
	cartoonRuneSlice := []rune(cartoon)
	for i := 0; i < len(cartoonRuneSlice); i++ {
		fmt.Printf("%c\r\n", cartoonRuneSlice[i])
	}
}
