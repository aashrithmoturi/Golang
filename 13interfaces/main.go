package main

import "fmt"

const PI = 3.14

type square struct {
	side float64
}
type circle struct {
	radius float64
}
type shape interface {
	area() float64
	circum() float64
}

// square method
func (s square) area() float64 {
	return s.side * s.side
}

func (s square) circum() float64 {
	return 4 * s.side
}

func (c circle) area() float64 {
	return PI * c.radius * c.radius
}

func (c circle) circum() float64 {
	return 2 * PI * c.radius
}

func printinfo(s shape) {
	fmt.Println(s.area())
	fmt.Println(s.circum())
}

func main() {
	shapes := []shape{
		square{1},
		square{2},
		circle{3},
		circle{4},
	}

	for _, val := range shapes {
		printinfo(val)
	}
}
