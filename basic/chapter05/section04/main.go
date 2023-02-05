package main

import "fmt"

func main() {
	var users = map[string]string{
		"鸣人": "下忍",
		"佐助": "叛忍",
		"小樱": "中忍",
	}

	// 遍历方式1
	for k, v := range users {
		fmt.Println(k, v)
	}

	// 遍历方式2
	for k := range users {
		fmt.Println(k, users[k])
	}

	// 是否存在
	d, ok := users["卡卡西"]
	if !ok {
		fmt.Println("not find")
	} else {
		fmt.Println("find", d)
	}

	// 删除k-v，即便删除不存在的k-v，也是不会报错的
	delete(users, "卡卡西")
	fmt.Println(users)
}
