package main

import (
	"fmt"
	"sync"
)

var x int

// func f(ch chan<- string) {
// 	ch <- "writing to channel"
// }

func increment(wg *sync.WaitGroup, ch chan int) {
	x += 1
	ch <- 6
	wg.Done()
}

func main() {
	// channel
	// ex
	// writing and reading from a unbuffered channel
	// ch := make(chan int)
	// go func(ch chan int) {
	// 	ch <- 8
	// 	fmt.Println("done writing to channel")
	// }(ch)

	// <-ch
	// fmt.Println("done reading from a channel")

	// for select
	// ch0 := make(chan string)
	// ch1 := make(chan string)
	// go func(ch chan string) {
	// 	ch <- "zero"
	// 	fmt.Println("done writing to channel")
	// }(ch0)

	// go func(ch chan string) {
	// 	ch <- "one"
	// 	fmt.Println("done writing to channel")
	// }(ch1)

	// for i := 0; i < 2; i++ {
	// 	select {
	// 	case msg1 := <-ch0:
	// 		fmt.Println("printing message from ch0 msg: ", msg1)
	// 	case msg2 := <-ch1:
	// 		fmt.Println("printing message from ch1 msg: ", msg2)
	// 	}
	// }

	// increment
	var wg sync.WaitGroup
	ch := make(chan int)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg, ch)
		<-ch
	}
	wg.Wait()
	fmt.Println("done incrementing, value of x ", x)

	// buffered channel
	// ch := make(chan string, 2)

	// ch <- "Hello"
	// ch <- "World"

	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	// directional channel
	// ch := make(<-chan int)

	// go channelD(ch)
	// fmt.Println("done reading from channel")

}
