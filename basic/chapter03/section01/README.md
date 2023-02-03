# 获取字符串长度

## 英文字符串长度

一个英文字符占 1 个字节

```go
func main() {
    nameEn := "Naruto"
    // 直接使用 len 求长度
    fmt.Println(len(nameEn))
}
```

## 中文字符串长度

一个中文字符在 utf-8 中占 3 个字节

```go
func main() {
    nameZh := "火影忍者"
    // 通过 rune 切片计算中文字符串的长度
    runeName := []rune(nameZh)
    fmt.Println(len(runeName))
}
```
