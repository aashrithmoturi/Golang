package main

import (
	"fmt"
	"sync"
)

func main() {
	myCh := make(chan int)
	// myCh <- 5
	// fmt.Println(<-myCh)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func (wg *sync.WaitGroup, ch chan int){
		
	}(wg, myCh)
	wg.Wait()
}
