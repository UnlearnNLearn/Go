package main

import "fmt"

type Person struct {
	Name string
	Age  uint
}

func main() {

	// for {
	// 	fmt.Println("1")
	// }

	slc := make([]int, 3, 5)

	fmt.Println(slc)

	// 2 way
	John := &Person{}
	Karan := new(Person)

	fmt.Println(John)
	fmt.Println(Karan)
	fmt.Printf("%T %T", John, Karan)

}
