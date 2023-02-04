# 切片的使用

## 概述

- Go 的切片("动态数组") 与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

## 切片的定义、赋值

```go
func main() {
    // 定义
    var classSeven []string // 定义切片初始长度为0，不可以使用索引方式添加！只能使用append方法动态添加

    // 动态添加元素
    classSeven = append(classSeven, "卡卡西")
    classSeven = append(classSeven, "鸣人")
    classSeven = append(classSeven, "佐助")
    classSeven = append(classSeven, "小樱")
}
```

## 切片的初始化方式

使用 string{} 方式

```go
func main() {
    // 初始化1
    var classEight []string = []string{"红", "雏田", "牙", "志乃"}

    // 初始化2
    classTen := []string{"阿斯玛", "鹿丸", "丁次", "井野"}
}
```

从数组直接创建

```go
func main() {
    kage := [5]string{"千手柱间", "千手扉间", "猿飞日斩", "波风水门", "纲手"}
    kageCopy1 := kage[:]
    kageCopy2 := kage[0:len(kage)] // [) 左闭右开
}
```

使用 make 函数

```go
func main() {
    threeNinjas := make([]string, 3)
    threeNinjas[0] = "大蛇丸"
    threeNinjas[1] = "自来也"
    threeNinjas[2] = "纲手"
}
```

## 访问切片元素

参考 python 切片用法，算是 python 的子集，不能有负数！

```go
func main() {
    // 支持
    fmt.Println(sliceData[0])
    fmt.Println(sliceData[1:4])
    fmt.Println(sliceData[2:])
    fmt.Println(sliceData[:6])
    fmt.Println(sliceData[:])
	
    // 不支持
    fmt.Println(sliceData[:-1])  // ❌ index -1 (constant of type int) must not be negative
}
```

## 添加切片元素

```go
func main() {
    var classSeven []string
    classSeven = append(classSeven, "卡卡西")
    classSeven = append(classSeven, "鸣人", "佐助", "小樱") // 支持多个值
    fmt.Println(classSeven) // [卡卡西 鸣人 佐助 小樱]

    // 切片的拼接
    classNew := []string{"大和", "鸣人", "小樱", "佐井"}
    classLast := []string{"凯", "宁次", "小李", "天天"}
    classData := append(classNew[1:], classLast[1:]...) // 三个点代表展开
    fmt.Println(classData) // [鸣人 小樱 佐井 宁次 小李 天天]
}
```

## 删除切片元素

删除切片其中的元素，虽然操作起来比较麻烦，但也只能这样拼接

```go
func main() {
    knife := []string{
        "断刀·斩首大刀",
        "大刀·鲛肌",
        "长刀·缝针",
        "钝刀·兜割", // 删除这个
        "爆刀·飞沫",
        "雷刀·牙",
        "双刀·鲆鲽",
    }
    newKnife := append(knife[:3], knife[4:]...)

    fmt.Println(knife)    // 原切片： [断刀·斩首大刀 大刀·鲛肌 长刀·缝针 钝刀·兜割 爆刀·飞沫 雷刀·牙 双刀·鲆鲽]
    fmt.Println(newKnife) // 删除后： [断刀·斩首大刀 大刀·鲛肌 长刀·缝针 爆刀·飞沫 雷刀·牙 双刀·鲆鲽]
}
```

## 切片的深浅拷贝

```go
func main() {
    // 拷贝
    monsters := []string{"蛤蟆文太", "蛤蟆吉", "蛤蟆龙"}

    // 浅拷贝
    monstersCopy := monsters[:]

    // 深拷贝
    var monstersCopyDeep = make([]string, len(monsters))
    copy(monstersCopyDeep, monsters)

    monsters[1] = "万蛇"
    monstersCopy[2] = "蛞蝓"
    fmt.Println(monsters)         // 原切片： [蛤蟆文太 万蛇 蛞蝓]
    fmt.Println(monstersCopy)     // 浅拷贝： [蛤蟆文太 万蛇 蛞蝓]
    fmt.Println(monstersCopyDeep) // 深拷贝： [蛤蟆文太 蛤蟆吉 蛤蟆龙]
}
```
