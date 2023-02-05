# Map 的使用

## 概述

- Go 中的 Map 是一个 key（索引）和 value（值）的无序集合，主要是查询方便，时间复杂度 `O(1)`！
- Go Map 中 value 的类型可以随意，但 key 的类型是有说法的！
- 注意：Map 是无序的，而且不能保证每次打印的顺序都是相同的！

## 重要提示

Go 内置的 Map 不是线程安全的！当有两个线程（协程） Goroutines，同时对一个 Map 进行操作，是会报错的！

如果在并发编程的情景下，想要使用 Map，需要使用：

```go
var syncMap sync.Map
```


## 哪些类型可作为 Map Key

> 可被比较的类型都可以作为 map key！基本是固定的值，可被 hash 的，才可以当作 key！

可作为 map key 的类型有：
- bool
- int
- float
- string
- pointer
- channel
- interface
- struct
- array

不能作为 map key 的类型有：
- slice
- map
- function

## Map 初始化与赋值

### Map 的初始化方式

```go
func main() {
    var users = map[string]string{
        "鸣人": "下忍",
        "佐助": "叛忍",
        "小樱": "中忍",
    }
    fmt.Println(users)
}
```

```go
func main() {
    // 使用make初始化
    users := make(map[string]string)

    // 设置k-v
    users["鸣人"] = "下忍"
    users["佐助"] = "叛忍"
    users["小樱"] = "中忍"
    fmt.Println(users)
}
```

### Map 容量

- 语法：`make(map[keytype]valuetype, cap)`

- 说明：
  - map 的初始容量 capacity 并非初始化必要的参数。
  - 和数组不同，map 可以根据新增的 key-value 对动态的伸缩，因此它不存在固定长度或最大限制。
  - 当 map 增长到容量上限的时候，如果再增加新的 key-value 对，map 的大小会自动加 1。
  - 出于性能的考虑，对于大的 map 或者会快速扩张的 map，即使只是大概知道容量，建议显性的标注。
  - map 是不可以使用 cap() 内置方法来查看容量的！

## Map 访问元素

```go
func main() {
    fmt.Println(users["卡卡西"])
}
```

## Map 的 for 循环

```go
func main() {
    // 遍历方式1
    for k, v := range users {
        fmt.Println(k, v)
    }

    // 遍历方式2
    for k := range users {
        fmt.Println(k, users[k])
    }
}
```

## 判断 Map 中是否存在元素

啰嗦写法

```go
func main() {
    d, ok := users["卡卡西"]
    if !ok {
        fmt.Println("not find")
    } else {
        fmt.Println("find", d)
    }
}
```

简洁写法

```go
func main() {
    if _, ok := users["卡卡西"]; !ok {
        fmt.Println("not find")
    } else {
        fmt.Println("find")
    }
}
```

## 删除 Map 元素

```go
func main() {
    // 即便删除不存在的k-v，也是不会报错的，可以放心使用
    delete(users, "卡卡西")
}
```
