package main

import (
	"fmt"
	"log"
)

func div(a, b int) int {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("%#v", err)
		}
	}()
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}

func main() {
	fmt.Println("Learning panic and recover")
	div(4, 2)
	div(4, 0)
	fmt.Println("now we now this program is recovered from panic")
}

// func panicker() {
// 	fmt.Println("about to panic")

// 	defer func() {
// 		if err := recover(); err != nil {
// 			log.Printf("%#v", err)
// 		}
// 	}()

// 	panic("test panic")
// }
