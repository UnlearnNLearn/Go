package main

import "fmt"

func main() {
	st := "we are learning about rune"

	streamBytes := []byte(st) // response body

	fmt.Println(streamBytes)

	var r rune = 'a'

	fmt.Println("value of r is", r)
	fmt.Printf("type of r is %T", r)

}
