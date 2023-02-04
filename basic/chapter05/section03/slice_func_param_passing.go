package main

import (
	"fmt"
)

// 内存地址变了，这是值传递，复制了一份数据，但是内部真正保存的元素是pointer类型，还是指向原来的位置，如改动值，则也会影响外部，效果如引用传递
// 发生扩容行为，则数据从指向原来的位置，完全分离开，不会影响外部，效果如值传递
func myfunc(data []int) {
	// [inner-before] slice: [1 2 3 4 5], idaddr: 0x1400011a048, length: 5, capacity: 5
	fmt.Printf("[inner-before] slice: %v, idaddr: %p, length: %d, capacity: %d\n", data, &data, len(data), cap(data))

	data[4] *= 10

	data = append(data, 66)

	// [inner-after] slice: [1 2 3 4 50 66], idaddr: 0x1400011a048, length: 6, capacity: 10
	fmt.Printf("[inner-after] slice: %v, idaddr: %p, length: %d, capacity: %d\n", data, &data, len(data), cap(data))
}

func main() {
	// 初始化
	sliceData := []int{1, 2, 3, 4, 5}

	// [outer-before] slice: [1 2 3 4 5], idaddr: 0x1400011a018, length: 5, capacity: 5
	fmt.Printf("[outer-before] slice: %v, idaddr: %p, length: %d, capacity: %d\n", sliceData, &sliceData, len(sliceData), cap(sliceData))

	myfunc(sliceData)

	// [outer-after] slice: [1 2 3 4 50], idaddr: 0x1400011a018, length: 5, capacity: 5
	fmt.Printf("[outer-after] slice: %v, idaddr: %p, length: %d, capacity: %d\n", sliceData, &sliceData, len(sliceData), cap(sliceData))
}
