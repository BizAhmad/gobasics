package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circ := Circle{0, 0, 5}
	rect := Rectangle{10, 5}

	fmt.Printf("Circle Area: %f \n", getArea(circ))
	fmt.Printf("Rectangle Area: %f \n", getArea(rect))
}
