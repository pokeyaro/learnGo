# Slice 与 Map 的细节

## 概述

- 高级类型如：map、slice、pointer、channel、func、interface 的零值，都是空 nil
- make 是内置函数，主要用于初始化 slice、map、channel 数据类型
- map 必须初始化才能够使用，但 slice 可以不初始化

## slice 的使用细节

> slice 不可以赋值超出它本身容量的数据，否则会报边界错误；\
> slice 即便超边界，也可使用扩容的方式，但该方式执行是非常低效的；\
> slice 仅声明类型，未初始化变量，零值为 nil，此时也是可以使用 append 方法进行扩容的；

### 仅声明类型（未初始化）

slice 零值

```go
func main() {
    // 仅声明变量，未初始化
    var s []string

    // s 为 nil
    if s == nil {
        fmt.Println("A nil slice")
    }
}
```

修改元素

```go
func main() {
    // 仅声明变量，未初始化
    var s []string
    fmt.Printf("val: %v, len: %d, cap: %d\n", s, len(s), cap(s)) // val: [], len: 0, cap: 0

    // 设置值（没有第0个元素，报错）
    s[0] = "a" // panic: runtime error: index out of range [0] with length 0
}
```

添加元素

```go
func main() {
    // 仅声明变量，未初始化
    var s []string
	fmt.Printf("val: %v, len: %d, cap: %d\n", s, len(s), cap(s)) // val: [], len: 0, cap: 0
    
    // 动态扩容（在最后添加元素）
    s = append(s, "a")
    fmt.Printf("val: %v, len: %d, cap: %d\n", s, len(s), cap(s)) // val: ["a"], len: 1, cap: 1
}
```

### 声明类型 + 初始化变量

修改元素

```go
func main() {
    // 声明变量且进行初始化（赋零值）
    s := make([]string, 1, 1)
    fmt.Printf("val: %v, len: %d, cap: %d\n", s, len(s), cap(s)) // val: [""], len: 1, cap: 1
    
    // 设置值（修改为新值，替换零值）
    s[0] = "a"
    fmt.Printf("val: %v, len: %d, cap: %d\n", s, len(s), cap(s)) // val: ["a"], len: 1, cap: 1
}
```

添加元素

```go
func main() {
    // 声明变量且进行初始化（赋零值）
    s := make([]string, 1, 1)
    fmt.Printf("val: %v, len: %d, cap: %d\n", s, len(s), cap(s)) // val: [""], len: 1, cap: 1
    
    // 动态扩容（不修改原值，在最后添加元素）
    s = append(s, "a")
    fmt.Printf("val: %v, len: %d, cap: %d\n", s, len(s), cap(s)) // val: ["", "a"], len: 2, cap: 2
}
```


## Map 的使用细节

> map 不初始化，在 nil 上直接加值，是不被允许的！

### 仅声明类型（未初始化）

map 零值

```go
func main() {
    // 仅声明变量，未初始化
    var m map[string]string

    // m 为 nil
    if m == nil {
        fmt.Println("A nil map")
    }
}
```

设置 k-v

```go
func main() {
    // 仅声明变量，未初始化
    var m map[string]string
    fmt.Printf("val: %v, len: %d\n", m, len(m)) // val: map[], len: 0

    // 设置一对k-v键值（报错）
    m["0"] = "a" // panic: assignment to entry in nil map
}
```

### 声明类型 + 初始化变量

方式1

```go
func main() {
    // 声明变量且进行初始化
    var m = map[string]string{}
    fmt.Printf("val: %v, len: %d\n", m, len(m)) // val: map[], len: 0

    // 设置k-v
    m["0"] = "a"
    m["1"] = "b"
    fmt.Printf("val: %v, len: %d\n", m, len(m)) // val: map[0:a 1:b], len: 2
}
```

方式2

```go
func main() {
    // 声明变量且进行初始化
    m := make(map[string]string, 1)
    fmt.Printf("val: %v, len: %d\n", m, len(m)) // val: map[], len: 0

    // 设置k-v
    m["0"] = "a"
    m["1"] = "b"
    fmt.Printf("val: %v, len: %d\n", m, len(m)) // val: map[0:a 1:b], len: 2
}
```
