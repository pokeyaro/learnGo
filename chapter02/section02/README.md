# 类型转换和格式化

## 数字之间的转换

```go
func main() {
    // 各种 int 类型之间的相互转换
    var i int8 = 12
    var x1 uint8 = uint8(i)
    fmt.Printf("%T, %v\n", x1, x1) // uint8, 12
	
    // 浮点型转整型
    var f float32 = 3.14
    var x2 = int32(f)
    fmt.Printf("%T, %v\n", x2, x2) // int32, 3

    // 整型转浮点型
    var n int8 = 50
    var x3 = float64(n)
    fmt.Printf("%T, %v\n", x3, x3) // float64, 50
}
```

## 类型别名

```go
func main() {
    type IT int
    var a IT = 8
    fmt.Printf("%T, %v\n", a, a) // main.IT, 8

    // IT 与 int 是不同的两个类型，不能直接运算或比较，会报错！这也正体现了go的强类型语言特性
    var b int = 8
    fmt.Println(int(a) - b)
}
```

## 字符串和整数的互转

通常使用 Atoi 和 Itoa 足够了

```go
import (
    "fmt"
    "strconv"
)

func main() {
    // 字符串转数字
    var intStr = "66"
    res1, err := strconv.Atoi(intStr)
    if err != nil {
        return
    }
    fmt.Printf("%T, %v\n", res1, res1) // int, 66

    // 数字转字符串
    var intNum = 33
    res2 := strconv.Itoa(intNum)
    fmt.Printf("%T, %v\n", res2, res2) // string, "33"
}
```

## 字符串转基本数据类型

```go
func main() {
    // 字符串转浮点型
    toFloat, err := strconv.ParseFloat("3.1415", 64)
    if err != nil {
        return
    }
    fmt.Printf("%T, %v\n", toFloat, toFloat) // float64, 3.1415

    // 字符串转整型（通常用于进制间转换）
    toInt, err := strconv.ParseInt("12", 8, 64)
    if err != nil {
        return
    }
    fmt.Printf("%T, %v\n", toInt, toInt) // int64, 10

    // 字符串转布尔
    toBool, err := strconv.ParseBool("true") // 1, 0, t, f, true, false ...
    if err != nil {
        return
    }
    fmt.Printf("%T, %v\n", toBool, toBool) // bool, true
}
```

## 基本数据类型转字符串

字符串格式化：更加定制化，强调对输出格式的 format

```go
func main() {
    // 布尔转字符串
    boolStr := strconv.FormatBool(false)
    fmt.Println(boolStr) // "false"

    // 浮点型转字符串（不同格式）
    floatStr := strconv.FormatFloat(3.1415926, 'E', -1, 64)
    fmt.Println(floatStr) // "3.1415926E+00"
	
    // 整型转自字符串（常进制转换）
    intStr := strconv.FormatInt(-42, 16)
    fmt.Println(intStr) // "-2a"
}
```
