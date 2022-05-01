package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	side float64
}

func main() {
	demoTriangle := triangle{base: 1.2, height: 1.2}
	demoSquare := square{side: 2}

	printArea(demoTriangle)
	printArea(demoSquare)
}

//calculate area of triangle
func (t triangle) getArea() float64 {
	return (t.base * t.height * 0.5)
}

//calculate area of square
func (s square) getArea() float64 {
	return s.side * s.side
}

func printArea(s shape) {
	fmt.Println("Area=[", s.getArea(), "]")
}
