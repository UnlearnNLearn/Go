package main

import "fmt"

// [1, 2 ] -> 2 -> 1 // taking reference of stack
func main() {
	//ex1
	// fmt.Println("I am learning defer keyword")

	// defer fmt.Println("first defer")  //1
	// defer fmt.Println("second defer") //2

	// fmt.Println("Goodbye! from main")

	//ex 2
	// for i := 0; i < 5; i++ {
	// 	defer fmt.Println(i) // 0, 1, 2, 3, 4 / 4, 3, 2, 1, 0
	// }

	// ex 3

	var z int = 44
	defer fmt.Println(z)
	z = 23
}
