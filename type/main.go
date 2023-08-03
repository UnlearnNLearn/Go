package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// var a int
	// var b bool
	// arr := [3]int{1, 2, 3}
	// slc := make([]int, 3, 5)
	// fmt.Printf("a = %#v type of a : %T\n", a, a)
	// fmt.Printf("b = %#v type of a : %T\n", b, b)
	// fmt.Printf("a = %#v type of a : %T\n", arr, arr)
	// fmt.Printf("b = %#v type of a : %T\n", slc, slc)

	// John := &Person{
	// 	Name: "John Doe",
	// 	Age:  34,
	// }
	// fmt.Println("Name of the person is", John.Name, " age of the person is ", John.Age)
	// fmt.Printf("John = %#v type of a : %T\n", John, John)
	mp := make(map[string]int, 3)
	mp["a"] = 1
	mp["b"] = 2
	mp["c"] = 3
	mp["d"] = 1
	fmt.Println(mp)
	mp["a"] = 2
	delete(mp, "d")
	fmt.Println(mp)

}
