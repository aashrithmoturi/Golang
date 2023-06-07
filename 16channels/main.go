package main

import "fmt"

func main() {
	myCh := make(chan int)
	myCh <- 5
	fmt.Println(<-myCh)

}
