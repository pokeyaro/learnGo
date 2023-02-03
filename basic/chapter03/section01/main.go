package main

import "fmt"

func main() {
	// 长度计算
	nameEn := "Naruto"
	fmt.Println("英文长度", len(nameEn)) // 英文长度 6

	nameZh := "火影忍者"
	fmt.Println("字节：", len(nameZh)) // 字节： 12
	
	runeName := []rune(nameZh)
	fmt.Println("中文长度：", len(runeName)) // 中文长度： 4
}
