# if 条件判断句

## if 语句结构

if 语句的结构，像其他语言一样有以下三种

```go
func main() {
    if 布尔表达式 {    // 如果不是复合条件判断，不需要使用 ()
        // 逻辑...
    }
}
```

```go
func main() {
    if 布尔表达式 {
        // 逻辑1...
    } else {
        // 逻辑2...
    }
}
```

```go
func main() {
    if 布尔表达式1 {
        // 逻辑1...
    } else if 布尔表达式2 {
        // 逻辑2...
    } else {
        // 逻辑3...
    }
}
```

## if 语句变体

使用简洁变量赋值，与判断逻辑放在一行进行处理，该变量的作用范围被限制在 if...else 语句组合中

```go
func main() {
    if num := 6; num%2 == 0 {
        fmt.Println(num, "是偶数")
    } else {
        fmt.Println(num, "是奇数")
    }
}
```

可以将返回值与判断放在一行进行处理，而且返回值的作用范围被限制在 if...else 语句组合中。

```go
package main

import (
    "errors"
    "fmt"
)

func promise() error {
    return errors.New("error")
}

func main() {
    if err := promise(); err != nil {
        fmt.Println("error happen")
    }
    fmt.Println(err)   // 此处会报错 undefined: err
}
```

Tips:
```text
在编程中，变量在其实现了变量的功能后，作用范围越小，所造成的问题可能性就越小。
每一个变量代表一个状态，有状态的地方，状态就会被修改，
函数的局部变量只会影响一个函数的执行，
但全局变量可能会影响所有代码的执行状态，
因此限制变量的作用范围对代码的稳定性有很大帮助。
```
