package main

import (
	"fmt"
)

func main() {
	// if... 条件句
	isNinja := true
	if isNinja {
		fmt.Println("一个忍者")
	}

	// if...else... 条件句
	age := 18
	if age < 18 {
		fmt.Println("小孩")
	} else {
		fmt.Println("成年")
	}

	// if...else if...else... 条件句
	name := "Naruto"
	if name == "Sakura" {
		fmt.Println("春野樱")
	} else if name == "Sasuke" {
		fmt.Println("宇智波佐助")
	} else {
		fmt.Println("漩涡鸣人")
	}

	// if 条件句的嵌套
	hero, country := "悟空", "日本"
	if hero == "悟空" {
		if country == "日本" {
			fmt.Println("七龙珠")
		} else {
			fmt.Println("西游记")
		}
	}

	// 复合条件（可以使用括号扩住单个条件）
	height, weight := 175, 70
	if (height > 180) && (weight < 60) {
		fmt.Println("高瘦子")
	} else if (height < 170) && (weight > 80) {
		fmt.Println("矮胖子")
	} else {
		fmt.Println("匀称")
	}
}
