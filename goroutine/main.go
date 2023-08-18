package main

import (
	"fmt"
	"runtime"
)

func main() {
	// goroutine
	// goroutine with waitgroup
	// number of goroutines goroutine

	fmt.Println("printing from main ")
	// 1
	fmt.Printf("my system has %#v number of CPU threads", runtime.GOMAXPROCS(-1))
	//2
	// var a int = 5
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	fmt.Println("this will print a's value and take it from closure", a)
	// 	wg.Done()
	// }()

	//3
	//var wg sync.WaitGroup
	// wg.Wait()
	// var wg sync.WaitGroup
	// for i := 0; i < 5; i++ {
	// 	wg.Add(1) // 5 counter
	// 	go func(i int) {
	// 		fmt.Println(i)
	// 		wg.Done() // decrement the counter
	// 	}(i)
	// }
	// wg.Wait() // 0

	//4
	// go fmt.Println("printing from goroutine")
	// time.Sleep(time.Millisecond * 50)
}
