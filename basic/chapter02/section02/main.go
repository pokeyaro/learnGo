package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 各种int类型之间的转换
	var i int8 = 12
	var x1 uint8 = uint8(i)

	fmt.Printf("%T, %v\n", x1, x1) // uint8, 12

	var f float32 = 3.14
	var x2 = int32(f)

	fmt.Printf("%T, %v\n", x2, x2) // int32, 3

	var n int8 = 50
	var x3 = float64(n)

	fmt.Printf("%T, %v\n", x3, x3) // float64, 50

	// 类型别名
	type IT int
	var a IT = 8
	fmt.Printf("%T, %v\n", a, a) // main.IT, 8

	var b int = 8
	fmt.Println(int(a) - b) // IT 与 int 是不同的两个类型，不能直接运算或比较，会报错！这也正体现了go的强类型语言特性

	// 字符串转数字
	var intStr = "66"
	res1, err := strconv.Atoi(intStr)
	if err != nil {
		fmt.Println("convert error")
	}
	fmt.Printf("%T, %v\n", res1, res1) // int, 66

	// 数字转字符串
	var intNum = 33
	res2 := strconv.Itoa(intNum)
	fmt.Printf("%T, %v\n", res2, res2) // string, 33

	// 将字符串转换为基础数据类型
	toFloat, err := strconv.ParseFloat("3.1415", 64)
	if err != nil {
		return
	}
	fmt.Printf("%T, %v\n", toFloat, toFloat) // float64, 3.1415

	toInt, err := strconv.ParseInt("12", 8, 64)
	if err != nil {
		return
	}
	fmt.Printf("%T, %v\n", toInt, toInt) // int64, 10

	toBool, err := strconv.ParseBool("true") // 1/0,t/f,true/false...
	if err != nil {
		return
	}
	fmt.Printf("%T, %v\n", toBool, toBool) // bool, true

	// 将基本数据类型转字符串
	boolStr := strconv.FormatBool(false)
	fmt.Printf("%T, %v\n", boolStr, boolStr) // string, false

	floatStr := strconv.FormatFloat(3.1415926, 'E', -1, 64)
	fmt.Printf("%T, %v\n", floatStr, floatStr) // string, 3.1415926E+00

	fmt.Println(strconv.FormatInt(-42, 16)) // -2a
}
