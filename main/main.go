package main

import (
	"fmt"
	"mth"
)

func init() {
	fmt.Println("inside init")
}

func main() {
	fmt.Println("inside main")
	fmt.Println("Hello, World")

	fmt.Println("5 + 4 = ", mth.Add(5, 4))
}
