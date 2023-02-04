# Slice 底层存储原理

## 概述

- Go 中的切片到底是 `值类型` 还是 `引用类型`？

## 引例

Go 的 slice 在作为函数参数传递的时候，是值传递还是引用传递？

是值传递，为什么效果上又呈现了引用的效果，但又不完全是！

```go
package main

import "fmt"

func genjutsu(userList []string) {
    // 看起来像是引用类型！
    userList[5] = "宇智波·鼬"

    // 那为啥没有扩容增加呢？
    for i := 1; i <= 10; i++ {
        userList = append(userList, fmt.Sprintf("影分身%d", i))
    }
}

func main() {
    ninja := []string{"自来也", "纲手", "卡卡西", "鸣人", "小樱", "佐助", "香燐", "重吾", "水月", "大蛇丸"}

    team1 := ninja[2:6]
    team2 := ninja[5:9]

    fmt.Println(team1) // [卡卡西 鸣人 小樱 佐助]
    fmt.Println(team2) // [佐助 香燐 重吾 水月]

    genjutsu(ninja)

    fmt.Println(team1) // [卡卡西 鸣人 小樱 宇智波·鼬]
    fmt.Println(team2) // [宇智波·鼬 香燐 重吾 水月]
}
```

## 结论先行

Go 函数参数传递切片，本质是值传递！

- 会复制一份数据到函数内，但真正保存的数据位置指针，内外部是共用的！
  - 当切片无扩容行为，如仅修改某个值，效果看起来像是引用传递！
  - 当切片发生扩容行为，则数据指针会转移到新的空间，内外部分离开，效果看起来像是值传递！

## 底层实现

Go 的切片是一种特殊的 "动态数组"，以 `[]string` 为例，它只是个语法糖的写法，本质上 slice 底层是个 `struct` 结构体

```go
type slice struct {
    array unsafe.Pointer // 用来存储实际数据的数组指针，指向一块连续的内存，仅指明 head 元素位置
    len int              // 切片中元素的数量
    cap int              // array 数组的长度
}
```

## 图解原理

- 代码：
[main.go](https://github.com/pokeyaro/learnGo/tree/master/basic/chapter05/section03/main.go)

- 图片：
[ProcessOn](https://www.processon.com/view/63ddf3d88363e12bdd67b1a0)

- 动态扩容：
[slice_dynamic_expansion.go](https://github.com/pokeyaro/learnGo/tree/master/basic/chapter05/section03/slice_dynamic_expansion.go)

