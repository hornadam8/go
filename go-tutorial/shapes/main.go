package main

import "fmt"

func main() {
	s := square{
		sideLength: 5.5,
	}
	t := triangle{
		base:   3.4,
		height: 4.6,
	}
	printArea(s)
	printArea(t)
}

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
