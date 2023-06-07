package main

import (
	"fmt"
	"net/http"
	"sync"
	// "time"
)

//	func greeter(s string) {
//		for i := 0; i < 4; i++ {
//			fmt.Println(s)
//		}
//	}
var signals = []string{}
var mut sync.Mutex

func getStatusCode(wg *sync.WaitGroup, endpoint string) {
	defer wg.Done()
	response, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Println(response.StatusCode, endpoint)
	}
}

func main() {
	// go greeter("Hello")
	// time.Sleep(1 * time.Millisecond)
	// greeter("world")
	var wg sync.WaitGroup
	websites := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
	}
	for _, web := range websites {
		wg.Add(1)
		go getStatusCode(&wg, web)
	}
	wg.Wait()
	// fmt.Println(signals) //if this line is above wg.wait() we would get empty
}
