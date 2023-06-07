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

	//receive only
	//we cannot use close statements on receive channels
	go func(wg *sync.WaitGroup, ch <-chan int) {
		//we must have same number of <-mych as myCh<-

		//If i donot have any senders then fp prints zero which shouldnt be the case
		//because we might have a case where sender sends a zero in both
		//cases printing zero doesnt make sense
		val, isChannelOpen := <-myCh
		if isChannelOpen == false {
			return
		}
		fmt.Println(val)
		fmt.Println(<-myCh)
		wg.Done()
	}(wg, myCh)

	//send only added a small comment here
	go func(wg *sync.WaitGroup, ch chan<- int) {
		myCh <- 5
		myCh <- 6
		close(myCh)
		wg.Done()
	}(wg, myCh)
	wg.Wait()
}
