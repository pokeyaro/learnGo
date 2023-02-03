package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	name := "鸣人"
	age := 18
	birthplace := "木叶村"
	var married bool

	//+号拼接方式（性能一般）
	fmt.Println("忍者: " + name + "\n年龄: " + strconv.Itoa(age) + "\n出生地: " + birthplace + "\n是否已婚: " + strconv.FormatBool(married))

	// 常用，直接格式化输出（性能较差）
	fmt.Printf("忍者: %s\n年龄: %d\n出生地: %s\n是否已婚: %v\n", name, age, birthplace, married)

	// 常用，可将格式化内容存储成字符串（性能较差）
	msg := fmt.Sprintf("忍者: %s\n年龄: %d\n出生地: %s\n是否已婚: %v\n", name, age, birthplace, married)
	fmt.Println(msg)

	// 特定场景，高性能方式
	var builder strings.Builder
	builder.WriteString("忍者: ")
	builder.WriteString(name)
	builder.WriteString("\n年龄: ")
	builder.WriteString(strconv.Itoa(age))
	builder.WriteString("\n出生地: ")
	builder.WriteString(birthplace)
	builder.WriteString("\n是否已婚: ")
	builder.WriteString(strconv.FormatBool(married))
	re := builder.String()
	fmt.Println(re)
}
