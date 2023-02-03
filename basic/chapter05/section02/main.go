package main

import "fmt"

func main() {
	var classSeven []string
	classSeven = append(classSeven, "卡卡西")
	classSeven = append(classSeven, "鸣人", "佐助", "小樱") // 支持多个值

	fmt.Println(classSeven) // [卡卡西 鸣人 佐助 小樱]

	// 切片的拼接
	classNew := []string{"大和", "鸣人", "小樱", "佐井"}
	classLast := []string{"凯", "宁次", "小李", "天天"}

	classData := append(classNew[1:], classLast[1:]...) // 三个点代表展开
	fmt.Println(classData)

	// 删除其中的元素
	knife := []string{
		"断刀·斩首大刀",
		"大刀·鲛肌",
		"长刀·缝针",
		"钝刀·兜割", // 删除这个
		"爆刀·飞沫",
		"雷刀·牙",
		"双刀·鲆鲽",
	}
	fmt.Println(knife)

	newKnife := append(knife[:3], knife[4:]...)
	fmt.Println(newKnife)

	// 拷贝
	frogs := []string{"蛤蟆文太", "蛤蟆吉", "蛤蟆龙"}

	// 浅拷贝
	frogsCopy := frogs[:]

	// 深拷贝
	var frogsCopyDeep = make([]string, len(frogs))
	copy(frogsCopyDeep, frogs)

	frogs[1] = "万蛇"
	frogsCopy[2] = "蛞蝓"
	fmt.Println("原切片：", frogs)
	fmt.Println("浅拷贝：", frogsCopy)
	fmt.Println("深拷贝：", frogsCopyDeep)
}
