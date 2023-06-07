package main

import "fmt"

func main() {
	defer fmt.Println("11")
	defer fmt.Println("12")
	fmt.Println("13")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

//guess
// 13, 4, 3, 2, 1, 0, 12, 11
