# for range 语句

for 循环的 range 格式可以对 slice、map、数组、字符串、channel 等进行迭代循环

## for range 格式

```go
func main() {
    for key, value := range object {
        fmt.Println(key, value)
    }
}
```

```go
func main() {
    // 使用匿名变量来忽略 key
    for _, value := range object {
        fmt.Println(value)
    }
}
```

## 不同数据类型的 for range

在不同的数据类型中，for range 的 key 代表是不一样

| 类型      | Key       | Value           | 省略 Key 的返回值 |
|---------|-----------|-----------------|-------------|
| 字符串     | 字符串的索引    | 字符串对应索引的字符串的拷贝  | 索引   |
| 数组      | 数组的索引     | 索引对应的值的拷贝       | 索引  |
| 切片      | 切片的索引     | 索引对应的值的拷贝       | 索引  |
| map     | map 的 key | map 的 value 的拷贝 | map 的 value |
| channel | -         | channel 接收的数据   | -           |

## 遍历字符串

```go
func main() {
    // utf-8 字符串
    var cartoon string = "Naruto 火影忍者 🍃"
	
    // 方式1 推荐
    for _, v := range cartoon {
        fmt.Printf("%c\n", v)
    }

    // 方式2 不推荐
    cartoonRuneSlice := []rune(cartoon)
    for i := 0; i < len(cartoonRuneSlice); i++ {
        fmt.Printf("%c\n", cartoonRuneSlice[i])
    }
}
```
