# 变量的作用域

## 概述

- 在静态语言中，变量作用域是有严格要求的

## 作用域示例

go 中可以直接写 `{}` 代码块

```go
package main

import "fmt"

func main() {
    {
        var localName = "local"
        fmt.Println(localName)
    }
    fmt.Println(localName) // 报错
}
```
