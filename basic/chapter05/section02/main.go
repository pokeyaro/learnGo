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
}
