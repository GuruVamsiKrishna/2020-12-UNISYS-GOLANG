package main

import (
	"fmt"
	"go-unisys/utils"
)

type person struct {
	name string
	age  int
	city string
}

func (p *person) print() {
	fmt.Printf("Name    : %v\n", p.name)
	fmt.Printf("Age     : %v\n", p.age)
	fmt.Printf("City    : %v\n", p.city)
	utils.Line()
}

func main() {
	p1 := person{"Vinod", 47, "Bangalore"}
	// p2 := person{"Vinod", 47} // ERROR: too few values in person literal
	p2 := person{name: "John Doe", age: 55}

	p3 := p2  // value
	p4 := &p2 // pointer

	p3.city = "Dallas"     // does not affect p2
	p4.city = "New Jersey" // affects p2 as well

	p1.print()
	p2.print()
	p3.print()
	p4.print()

}
