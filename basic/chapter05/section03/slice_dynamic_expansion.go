package main

import "fmt"

// InSlice 判断 int 是否在 slice 中
func InSlice(origin []int, n int) bool {
	for _, elem := range origin {
		if elem == n {
			return true
		}
	}
	return false
}

// 观察 Go slice 切片动态扩容的容量增长趋势与规律
func main() {
	// 初始化
	sliceData := make([]int, 1) // [0]

	// 动态扩容的增长
	growth := make([]int, 0, 20) // []

	// 动态添加数据
	for i := 1; i <= len(sliceData); i++ {
		sliceData = append(sliceData, i)

		// 判断是否在集合中
		if InSlice(growth, cap(sliceData)) {
			continue
		} else {
			growth = append(growth, cap(sliceData))
		}

		// 跳出循环
		if len(sliceData) > 10000 {
			goto label
		}
	}

label:
	// 在大小为 512 之前，当 len==cap 的时候，会 x2 的扩容，2 4 8 16 32 64 128 256 512
	fmt.Println(growth) // [2 4 8 16 32 64 128 256 512 848 1280 1792 2560 3408 5120 7168 9216 12288 16384]
	fmt.Printf("length: %d, capacity: %d\n", len(sliceData), cap(sliceData)) // length: 12289, capacity: 16384
	return
}
