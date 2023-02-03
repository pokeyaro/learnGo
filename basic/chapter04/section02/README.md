# for 循环语句

> 在 Go 中，只有 for 循环，没有 while 关键字

## for 语句结构

Go 的 for 循环提供了 3 种形式

- `init`： 一般为赋值表达式，给控制变量赋初值；
- `condition`： 关系表达式或逻辑表达式，循环控制条件；
- `post`： 一般为赋值表达式，给控制变量增量或减量；

```go
func main() {
    // 和 C 语言的 for 一样：
    for init; condition; post {
        // 循环体...
    }
}
```

```go
func main() {
    // 和 C 的 while 一样：
    for condition {
        // 循环体...
    }
}
```

```go
func main() {
    // 和 C 的 for(;;) 一样：
    for { 
        // 循环体...
    }
}
```

## for 代码示例

```go
func main() {
    // i变量限定在循环体内
    for i := 1; i < 10; i++ {
        fmt.Println(i)
    }

    // 变体1
    var x int = 1
    for ; x < 10; x++ {
        fmt.Println(x)
    }

    // 变体2
    var y int = 1
    for y < 10 {
        fmt.Println(y)
        y++
    }

    // 死循环
    for {
        fmt.Println("ok")
    }
}
```

## break 和 continue

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    for round := 0; ; round++ {
        time.Sleep(1 * time.Second)
        if round > 10 {
            fmt.Println("exit")
            break // 跳出该循环
        }
        if round%2 == 0 {
            fmt.Println("skip")
            continue // 跳过当次循环
        }
        fmt.Println(round)
    }
}
```
