package main

import (
	"container/list"
	"fmt"
)

func main() {
	var userList list.List

	// 向尾部添加元素
	userList.PushBack("鸣人")
	userList.PushBack("佐助")
	userList.PushBack("小樱")

	// 向头部插入数据
	userList.PushBack("卡卡西")

	fmt.Println(userList) // {{0x140001121b0 0x14000112240 <nil> <nil>} 4}

	// 遍历打印值，正序
	for i := userList.Front(); i != nil; i = i.Next() {
		fmt.Printf("%v ", i.Value)
	}

	// 遍历打印值，逆序
	for i := userList.Back(); i != nil; i = i.Prev() {
		fmt.Printf("%v ", i.Value)
	}

	// 删除元素
	i := userList.Front()
	for ; i != nil; i = i.Next() {
		if i.Value.(string) == "佐助" {
			goto label
		}
	}
label:
	userList.Remove(i)
	for i := userList.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
