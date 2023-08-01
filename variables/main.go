package main

import "fmt"

type Entity string

func main() {
	// int
	// var a int = 25
	// string
	var st Entity = "we are learning how to declare and initialize variables"
	// var f float64 = 2.5
	// b := 3
	var a, b int = 2, 3
	const c = 4
	fmt.Printf("a = %#v b= %#v c=%#v st=%#v \n", a, b, c, st)
}
