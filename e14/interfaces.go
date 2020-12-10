package main

import (
	"fmt"
	"math"
)

type Triangle struct {
	Base, Height float64
}

func (t Triangle) GetArea() (area float64) {
	area = 0.5 * t.Base * t.Height
	return
}

type Circle struct {
	Radius float64
}

func (c Circle) GetArea() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Shape interface {
	GetArea() float64
}

// polymorphic function; works with different types of shapes as long as
// they have GetArea() function returning float64
func Measure(t Shape) {
	area := t.GetArea()
	fmt.Println("Area is", area)
}

func main() {
	fmt.Println("Interfaces")

	t1 := Triangle{12.34, 56.78}
	Measure(t1)

	c1 := Circle{12.34}
	Measure(c1)
}
