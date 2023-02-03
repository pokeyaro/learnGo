package main

import (
	"fmt"
	"strings"
)

func main() {
	originStr := "naruto.mingren@mail.com"

	// 是否包含xx子串
	v1 := strings.Contains(originStr, "mail") // true

	// 是否以xx子串开头
	v2 := strings.HasPrefix(originStr, "naruto") // true

	// 是否以xx子串结尾
	v3 := strings.HasSuffix(originStr, "com") // true

	// xx子串出现的次数
	v4 := strings.Count(originStr, "o") // 2

	// 字符串的分割
	v5 := strings.Split(originStr, "@") // [naruto.mingren mail.com]

	// 查找xx子串出现的位置
	v6 := strings.Index(originStr, "ming") // 7

	// 子串替换
	v7 := strings.Replace(originStr, "mingren", "zuozhu", 1)   // 替换指定次数，若为 -1，也相当于全部替换
	v8 := strings.ReplaceAll(originStr, "mingren", "xiaoying") // 全部替换

	// 大小写转换
	v9 := strings.ToLower("Hello Naruto")  // hello naruto
	v10 := strings.ToUpper("Hello Naruto") // HELLO NARUTO

	// 去掉特殊字符
	v11 := strings.Trim("###$$ Naruto $$###", "$# ")
	v12 := strings.TrimLeft("  \tNaruto", " \t")
	v13 := strings.TrimRight("Naruto\t  ", " \t")

	fmt.Println(v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13)
}
