# List 的使用

## 与 Slice 的区别

- Slice：动态数组，在发生扩容时，会频繁的分配空间，数据拷贝，Slice 的底层一定是连续的存储空间！

- List：链表结构，底层可以不是连续的！在每个元素中都需要有一部分空间存储下一个位置的指针，这样就造成了空间浪费！ 

- List 虽然更灵活，但与 slice 相比，由于数据结构不同，list 需要不停的根据指针才能够知道下一个元素位置在哪，所以两者性能差异很大。

## 适用场景

- 如果是在频繁插入/删除元素的场景，我们或许会优先选择 list，但别忘了 list 元素的查询是很慢的！

- 在绝大多数场景下，使用 slice 会是更优的选择！

## list 初始化

```go
import (
    "container/list"
)

func main() {
    // 方式1
    var userList list.List

    // 方式2
    userList := list.New()
}
```

## list 添加元素

```go
func main() {
    // 向尾部添加元素
    userList.PushBack("鸣人")
    userList.PushBack("佐助")
    userList.PushBack("小樱")

    // 向头部插入数据
    userList.PushBack("卡卡西")

    // 打印
    fmt.Println(userList) // {{0x140001121b0 0x14000112240 <nil> <nil>} 4}
}
```

## list 的遍历

```go
func main() {
    // 遍历打印值，正序
    for i := userList.Front(); i != nil; i = i.Next() {
        fmt.Printf("%v ", i.Value)
    }

    // 遍历打印值，逆序
    for i := userList.Back(); i != nil; i = i.Prev() {
        fmt.Printf("%v ", i.Value)
    }
}
```

## list 插入元素

```go
func main() {
    // 找到标志元素位置
    i := userList.Front()
    for; i != nil; i = i.Next() {
        if i.Value.(string) == "鸣人" {
            break
        }
    }

    // 在标志位置之前插入
    userList.InsertBefore("好色仙人", i)

    // 在标志位置之后插入
    userList.InsertAfter("蛤蟆文太", i)
}
```

## list 删除元素

```go
func main() {
    // 查找标记元素
    i := userList.Front()
    for; i != nil; i = i.Next() {
        if i.Value.(string) == "佐助" {
            break
        }
    }

    // 删除标记元素
    userList.Remove(i)
}
```

