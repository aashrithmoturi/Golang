package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghgwuyej"

func main() {
	fmt.Println("Handling url")
	fmt.Println(myurl)

	// parsing the url
	res, _ := url.Parse(myurl)
	// fmt.Println(res.Scheme)
	// fmt.Println(res.User)
	// fmt.Println(res.Host)
	// fmt.Println(res.Port())
	// fmt.Println(res.Path)
	// fmt.Println(res.RawQuery)

	qparams := res.Query()
	fmt.Printf("%T\n", qparams)
	// fmt.Println(qparams)
	for key, value := range qparams {
		fmt.Println(key, value)
	}

	//& is important
	partOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/qjhwgf",
		RawQuery: "user=ash",
	}

	//.string() parse it
	fmt.Println(partOfUrl.String())

}
