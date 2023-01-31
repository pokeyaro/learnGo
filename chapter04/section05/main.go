package main

import "fmt"

func main() {
	var grade byte = 'C'
	var level string

	switch grade {
	case 'S':
		level = "影"
	case 'A':
		level = "上忍"
	case 'B', 'C':
		level = "中下忍"
	default:
		level = "未毕业"
	}

	switch {
	case level == "影":
		fmt.Printf("首脑领袖\n")
	case level == "上忍" || level == "中下忍":
		fmt.Printf("普通忍者\n")
	default:
		fmt.Printf("村民\n")
	}
}
