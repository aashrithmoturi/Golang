package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.reddit.com"

func main() {

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T\n", response)
	//its caller's responsibility to close the connection
	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)

}
