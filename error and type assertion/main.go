package main

import "fmt"

// type assertion
// switch
// rune - problem

func main() {
	var x interface{} = "foo"

	s := x.(string)
	fmt.Println("checking the value of x ", s)

	t, ok := x.(string)
	if !ok {
		fmt.Println("ERROR type asserting x to string")
	}
	fmt.Println("checking the value of x ", t)

	i, ok := x.(int) // , ok syntax
	if !ok {
		fmt.Println("ERROR type asserting x to int")
	}
	fmt.Println("checking the value of x ", i)

}
