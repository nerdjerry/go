package main

import (
	"fmt"
)

type triangle struct {
	base   float64
	height float64
}
type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println("Area of shape is: ", s.getArea())
}
func main() {
	t := triangle{
		base:   0.4,
		height: 0.6,
	}
	s := square{
		sideLength: 1.2,
	}
	printArea(t)
	printArea(s)
}
