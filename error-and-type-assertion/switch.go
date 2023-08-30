package main

import "fmt"

func main() {
	var x interface{} = 23

	switch x.(type) {
	case nil:
		fmt.Println("type of x is nil")
	case string, bool:
		fmt.Println("type of x is either string or bool")
	case int:
		fmt.Println("type of x is int")
	default:
		fmt.Println("not able to determine the type of x")

	}
}
