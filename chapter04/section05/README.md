# switch 语句

## 概述

- switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上至下逐一测试，直到匹配为止。

## switch 语法格式

```go
func main() {
    switch variable {
    case res1:
        ...
    case res2:
        ...
    default:
        ...
    }
}
```
```go
func main() {
    switch {
    case condition1:
        ...
    case condition2:
        ...
    default:
        ...
    }
}
```

## case 表达式

### 一分支多值

不同的 case 表达式使用逗号分隔

```go
switch {
case "Mum", "Dad":
    fmt.Println("We are 伐木累")
}
```

### 分支表达式

```go
switch {
case n > 10 && n < 20:
    fmt.Println("10 < n < 20")
}
```

## Type Switch

switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。

```go
func main() {
    var x interface{}

    switch i := x.(type) {
    case nil:
        fmt.Printf(" x 的类型 :%T", i)
    case bool:
        fmt.Printf("x 是 bool 型")
    case int64:
        fmt.Printf("x 是 int64 型")
    case float64:
        fmt.Printf("x 是 float64 型")
    case string:
        fmt.Printf("x 是 string 型")
    case func(int) int:
        fmt.Printf("x 是 func(int) int 型")
    default:
        fmt.Printf("未知型")
    }
}
```
