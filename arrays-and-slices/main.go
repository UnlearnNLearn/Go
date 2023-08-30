package main

import "fmt"

func manipulateArray(arr [3]int) {
	arr[0] = 3
}

func manipulateSlice(arr []int) {
	arr[0] = 3
}

func main() {
	arr := [3]int{1, 2, 3}
	slc := make([]int, 3, 5)
	slc = arr[:]
	slc = append(slc, 3, 4, 5, 6)
	slc[0] = 5
	for i, a := range slc {
		fmt.Println("index =", i, "value =", a)
	}
	fmt.Println("index 0 ", arr[0])
}
