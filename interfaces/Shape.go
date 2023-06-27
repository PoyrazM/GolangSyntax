package interfaces

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

type triangle struct {
	base   float64
	height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (t triangle) area() float64 {
	return t.base * (t.height / 2)
}

func calculate(s shape) {
	fmt.Println("The area of the shape :", s.area())
}

func CalculateArea() {
	r := rectangle{width: 10, height: 6}
	calculate(r)

	c := circle{radius: 10}
	calculate(c)

	t := triangle{base: 10, height: 2}
	calculate(t)
}
