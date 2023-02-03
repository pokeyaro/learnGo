# iota的使用细节

## 概述

- iota，是一个特殊的常量，可以认为是一个可被编译器修改的常量 
- iota 的内部有个计数器，它从 0 开始，只要遇到 iota 就 +1

## iota使用

iota 能够简化 const 类型的定义

```go
package main

func main() {
    const (
        ERR01 = iota
        ERR02
        ERR03
    )
}
```

如果中断了 iota，那么必须显式的恢复，后续会自动递增

```go
package main

func main() {
    const (
        D1 = iota + 1 // 自增类型默认是int类型
        D2
        S1 = "abc"    // iota内部仍然会增加计数器
        S2
        D4 = iota     // 显式恢复，赋值当前的iota数值
        D5
    )
}
```

每次出现 const 的时候，iota 初始值都归 0

```go
package main

func main() {
    const (
        V1 = iota
        V2
    )
    const (
        K1 = iota
    )
}
```
