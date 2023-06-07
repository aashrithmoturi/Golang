package main

import (
	"fmt"
	"sort"
)

func main() {
	var array = []string{"A", "B"}
	// fmt.Printf("%T", array)

	array = append(array, "a", "aa", "aaa")
	fmt.Println(array)

	// array = append(array[1:])
	// fmt.Println(array)

	//while using the end-range we must
	//make sure that it is non-inclusive

	// array = append(array[1:3])
	// fmt.Println(array)

	// array = append(array[:3])
	// fmt.Println(array)

	highScores := make([]int, 4)
	highScores[0] = 1
	highScores[1] = 2
	highScores = append(highScores, 12, 13)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)

	//remove an element from slice based on index
	var courses = []int{1, 2, 3, 4, 5, 6}
	courses = append(courses[:2], courses[3:]...)
	fmt.Println(courses)

}
