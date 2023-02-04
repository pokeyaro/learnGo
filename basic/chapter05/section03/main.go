package main

import "fmt"

const (
	// LIMIT_NUM = 2 // 不会扩容
	LIMIT_NUM = 3 // 会扩容
)

func operate(s []int) {
	fmt.Println("-----进入函数体-----")
	fmt.Printf("[sliceData] [before] 值：%v, 地址：%p, 长度：%d, 容量：%d\n", s, &s, len(s), cap(s))

	fmt.Println("-----进行切片（内部）-----")
	sliceInner := s[3:8]
	fmt.Printf("[sliceInner] [before] 值：%v, 地址：%p, 长度：%d, 容量：%d\n", sliceInner, &sliceInner, len(sliceInner), cap(sliceInner))

	fmt.Println("-----进行修改（内部）-----")
	sliceInner[0] *= 100

	for i := 1; i <= LIMIT_NUM; i++ {
		sliceInner = append(sliceInner, i*10)
	}

	fmt.Printf("[sliceInner] [after] 值：%v, 地址：%p, 长度：%d, 容量：%d\n", sliceInner, &sliceInner, len(sliceInner), cap(sliceInner))

	fmt.Printf("[sliceData] [after] 值：%v, 地址：%p, 长度：%d, 容量：%d\n", s, &s, len(s), cap(s))
	fmt.Println("-----离开函数体-----")
}

func main() {
	sliceData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("-----执行函数前-----")
	fmt.Printf("[sliceData] [before] 值：%v, 地址：%p, 长度：%d, 容量：%d\n", sliceData, &sliceData, len(sliceData), cap(sliceData))

	fmt.Println("-----进行切片（外部）-----")
	sliceOuter := sliceData[2:5]
	fmt.Printf("[sliceOuter] [before] 值：%v, 地址：%p, 长度：%d, 容量：%d\n", sliceOuter, &sliceOuter, len(sliceOuter), cap(sliceOuter))

	fmt.Println("-----执行函数-----")
	operate(sliceData)

	fmt.Println("-----执行函数后-----")
	fmt.Printf("[sliceData] [before] 值：%v, 地址：%p, 长度：%d, 容量：%d\n", sliceData, &sliceData, len(sliceData), cap(sliceData))

	fmt.Println("-----进行修改（外部）-----")
	sliceOuter[2] *= 100
	fmt.Printf("[sliceOuter] [after] 值：%v, 地址：%p, 长度：%d, 容量：%d\n", sliceOuter, &sliceOuter, len(sliceOuter), cap(sliceOuter))

	fmt.Println("-----最终-----")
	fmt.Printf("[sliceData] [after] 值：%v, 地址：%p, 长度：%d, 容量：%d\n", sliceData, &sliceData, len(sliceData), cap(sliceData))
}
