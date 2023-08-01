package main

import "fmt"

func init() {
	fmt.Println("inside init")
}

func sum(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("inside main")
	fmt.Println("Hello, World")

	fmt.Println("5 + 4 = ", sum(5, 4))
}
