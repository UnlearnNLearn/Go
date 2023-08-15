package main

import "fmt"

type Receiver int

func (r Receiver) funcAcceptingReceiver() {
	fmt.Println("this function is called using receiver")
}

func main() {
	fmt.Println("learning method and receiver")
	var r Receiver
	r.funcAcceptingReceiver()
}
