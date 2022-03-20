package main

import "fmt"

func main() {
	t := triangle{
		height: 4,
		base:   3,
	}

	s := square{
		sideLenght: 3,
	}

	printArea(s)
	printArea(t)
}

type square struct {
	sideLenght float64
}

type triangle struct {
	height float64
	base   float64
}

type shape interface {
	getArea() float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLenght * s.sideLenght
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
