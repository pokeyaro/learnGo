package main

import "fmt"

// 全局变量和局部变量
var (
	name = "Naruto"
	age  = 18
	male bool
)

func main() {

	// 定义单个变量
	var classNum int
	classNum = 7

	var classZh string = "第七班"

	classEn := "seventh class"

	// 使用变量
	fmt.Println(classNum, classZh, classEn)

	// 定义多个变量
	var member1, member2, member3 = "鸣人", "佐助", "小樱"

	fmt.Println(member1, member2, member3)

	// 多变量定义 不同类型
	name, age, male := "鸣人", 8, true

	fmt.Println(name, age, male)

	var user string // ""
	var passwd int  // 0
	var active bool // false
	fmt.Println(user, passwd, active)
}
