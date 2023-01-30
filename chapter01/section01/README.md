# 如何定义变量

## 概述

- go 是静态语言，静态语言和动态语言相比，变量差异较大：

1. 变量必须先定义，后使用
2. 变量声明类型与赋值类型一致
3. 类型定义后，不允许再改变
4. 变量名不能冲突

## 定义单个变量

多种定义变量的方式

```go
package main

import "fmt"

func main() {
    // 方式一
    var classNum int
    classNum = 7
	
    // 方式二 
    var classZh string = "第七班"

    // 方式三 
    classEn := "seventh class"

    // 使用变量
    fmt.Println(classNum, classZh, classEn)
}
```

go 语言中定义变量后必须使用，否则会报错 `“Unused variable 'value'”`

## 定义多个变量

多变量定义，可以是相同类型，也可以是不同类型

```go
func main() {
    var member1, member2, member3 = "鸣人", "佐助", "小樱"
    name, age, male := "鸣人", 8, true
}
```

## 全局变量与局部变量

简洁变量定义不能作用到全局，否则会报错

```go
package main

import "fmt"

// 全局变量
var (
    name = "naruto"
    age  = 10
    male bool
)

func main() { 
    // 局部变量优先级高于全局变量
    name, age, male := "鸣人", 18, true
    fmt.Println(name, age, male)
}

```

## 变量的零值

go 语言变量是有零值的，如下：

```go
func main() {
    var user string // ""
    var passwd int  // 0 
    var active bool // false
}
```