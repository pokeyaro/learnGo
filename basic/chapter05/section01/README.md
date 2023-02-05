# 数组的使用

## 概述

- Go 的数组是具有相同唯一类型的一组已编号且长度固定的数据项序列。这种类型可以是任意的原始类型例如整形、字符串或者自定义类型。

- 数组的大小是类型的一部分。[5]int 和 [6]int 是不同的类型。因此，数组不能被调整大小。

- Go 的数组是值类型，而不是引用类型。这意味着当它们被分配给一个新变量时，将把原始数组的副本分配给新变量。如果对新变量进行了更改，则不会在原始数组中反映。

## 数组的定义、赋值

```go
func main() {
    // 定义
    var classSeven [4]string

    // 赋值
    classSeven[0] = "卡卡西"
    classSeven[1] = "鸣人"
    classSeven[2] = "佐助"
    classSeven[3] = "小樱"
}
```

## 初始化数组

- 初始化数组中 {} 中的元素个数不能大于 [] 中的数字。
- 如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小。

```go
func main() {
    // 方式1
    var classEight [4]string = [4]string{"红", "雏田", "牙", "志乃"}

    // 方式2
    classTen := [...]string{"阿斯玛", "鹿丸", "丁次", "井野"}

    // 方式3：初始化部分
    kage := [5]string{0: "千手柱间", 3: "波风水门"}
}
```

## 访问数组元素

- 数组元素可以通过索引（位置）来读取（或者修改），索引从 0 开始，第一个元素索引为 0。

```go
func main() {
    // 读取
    fmt.Println(classSeven)     // [卡卡西 鸣人 佐助 小樱]
    fmt.Println(classSeven[0])  // 卡卡西
    fmt.Println(classSeven[:2]) // [卡卡西 鸣人]
}
```

## 遍历数组

```go
func main() {
    // 方式1
    for index, user := range classSeven {
        fmt.Println(index, user)
    }

    // 方式2
    for i := 0; i < len(classSeven); i++ {
        fmt.Println(classTen[i])
    }
}
```

## 多维数组

```go
func main() {
    // 定义一个 '三行四列' 的多维数组
    var classInfo [3][4]string

    classInfo[0] = [4]string{"卡卡西", "鸣人", "佐助", "小樱"}
    classInfo[1] = [4]string{"红", "雏田", "牙", "志乃"}
    classInfo[2] = [4]string{"阿斯玛", "鹿丸", "丁次", "井野"}

    for i := 0; i < len(classInfo); i++ {
        for j := 0; j < len(classInfo[i]); j++ {
            fmt.Printf("%v ", classInfo[i][j])
        }
        fmt.Println()
    }
}
```
