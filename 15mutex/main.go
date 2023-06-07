package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("This is all about mutex!!!")
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	wg.Add(3)

	var arr = []int{}

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		m.Lock()
		arr = append(arr, 1)
		m.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		m.Lock()
		arr = append(arr, 2)
		m.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		// fmt.Println(arr)
		m.Lock()
		arr = append(arr, 3)
		m.Unlock()
		wg.Done()
	}(wg, mut)

	//fmt.Println(arr) adding it here gives race condition error
	wg.Wait()
	fmt.Println(arr)

}
