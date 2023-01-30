package main

import "fmt"

var name = "naruto"

func main() {
	var age int
	if name == "naruto" {
		age := 18
		fmt.Println("inner naruto", age) // inner naruto 18
	} else {
		age = 20
		fmt.Println("inner other", age)
	}
	fmt.Println("outer", age) // outer 0
}
