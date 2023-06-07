package main

import "fmt"

func main() {
	var num [3]int
	num[0] = 1
	num[1] = 2
	num[2] = 3
	fmt.Println(num)
	fmt.Println(len(num))

	var a = [3]string{"a", "b", "d"}
	fmt.Println(a)
}
