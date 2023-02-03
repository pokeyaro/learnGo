# goto 关键字

## 概述

- Go 语言的 goto 语句可以无条件地转移到过程中指定的行。
- goto 语句通常与条件语句配合使用。可用来实现条件转移，构成循环，跳出循环体等功能。
- 但是，在结构化程序设计中一般不主张使用 goto 语句，以免造成程序流程的混乱，使理解和调试程序都产生困难。

## goto 语法格式

```go
func main() {
    goto label;
    ...
    label: statement;
}
```

## 使用 goto 跳出多层循环

```go
func main() {
    for x := 0; x < 10; x++ {
        for y := 0; y < 10; y++ {
            if y == 2 {
                // 跳转到标签
                goto label
            }
        }
    }
    // 手动返回, 避免执行进入标签
    return
// 标签
label:
    fmt.Println("here")
}
```

## 使用 goto 集中处理错误

```go
func main() {
    err := firstCheckError()
    if err != nil {
        goto onExit
    }
    err = secondCheckError()
    if err != nil {
        goto onExit
    }
    fmt.Println("done")
    return
onExit:
    fmt.Println(err)
    exitProcess()
}
```
