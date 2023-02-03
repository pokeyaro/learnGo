package main

import "fmt"

func main() {
	// 数组的定义
	var classSeven [4]string

	// 赋值
	classSeven[0] = "卡卡西"
	classSeven[1] = "鸣人"
	classSeven[2] = "佐助"
	classSeven[3] = "小樱"

	fmt.Println(classSeven[0])

	// 初始化
	var classEight [4]string = [4]string{"红", "雏田", "牙", "志乃"}

	fmt.Println(classEight)

	// 简化
	classTen := [...]string{"阿斯玛", "鹿丸", "丁次", "井野"}

	// 遍历数组
	for k, v := range classTen {
		fmt.Println(k, v)
	}

	for i := 0; i < len(classTen); i++ {
		fmt.Println(classTen[i])
	}

	// 初始化其中某几个
	Kage := [5]string{0: "千手柱间", 3: "波风水门"}
	fmt.Println(Kage, len(Kage))
}
