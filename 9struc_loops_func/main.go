package main

import "fmt"

func main() {
	//There is no inheritance in golang
	ash := user{"Ash", "a@gmail.com", true, 21}
	fmt.Println(ash)
	fmt.Printf("%+v\n", ash)

	days := make([]string, 0)
	days = append(days, "A", "B", "C")

	// //type-1
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// //type-2
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Printf("%v, %v\n", index, day)
	}
	fmt.Println("")

	i := 1
	for i < 10 {
		if i == 4 {
			i++
			continue
		}
		if i == 9 {
			break
		}
		fmt.Println(i)
		i++
	}
	fmt.Println("")

	//!!!!!!!   Functions   !!!!!!!!!!
	fmt.Println(f(10, 2, 3))

}

func f(arr ...int) int {
	sum := 0
	for i := range arr {
		sum += arr[i]
	}
	return sum
}

type user struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
