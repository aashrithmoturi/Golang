package main

import "fmt"

func main() {
	num := 2
	ptr := &num
	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = *ptr + 1
	fmt.Println(num)
}
