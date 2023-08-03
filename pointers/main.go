package main

import "fmt"

//variadic function
func sum(a ...int) {
	total := 0
	for _, v := range a {
		total += v
	}
	fmt.Println("total of numbers provided is", total)
}

func main() {
	// a[5] <- b
	var a int = 5
	var b *int = &a

	fmt.Printf("a = %#v type of a : %T\n", a, a)
	fmt.Printf("b = %#v type of b : %T\n", *b, b)

	//calling variadic function

	sum(1, 2, 3, 4)
	sum(5, 4)

	//anynomous function
	func() {
		fmt.Println("inside anynomous function")
	}()
}
