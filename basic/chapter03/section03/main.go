package main

import (
	"fmt"
)

func main() {
	var byteSlice []byte = []byte{'A', 'B', 'C'}

	fmt.Printf("%T\n", byteSlice) // []uint8

	fmt.Printf("%v\n", byteSlice) // [65 66 67]

	fmt.Printf("%#v\n", byteSlice) // []byte{0x41, 0x42, 0x43}

	fmt.Printf("%c\n", byteSlice) // [A B C]

	fmt.Printf("%s\n", byteSlice) // ABC
}
