package main

import (
	"fmt"
	"go-unisys/utils"
)

// func <name>(arg list) <ret-type> { function-body }
func add(a int, b int) int {
	return a + b
}

func div(a int, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

func sum(data ...int) int {
	// data is a slice of int
	// fmt.Printf("%v %T\n", data, data)
	s := 0
	for _, n := range data {
		s += n
	}
	return s
}

type person struct {
	gender, fname, lname string
}

func (this person) fullname() string {
	title := "Mr."
	if this.gender != "Male" {
		title = "Ms."
	}
	return title + this.fname + " " + this.lname
}

func main() {

	p1 := person{"Male", "Vinod", "Kumar"}
	p2 := person{"Female", "Jane", "Doe"}

	fmt.Println(p1.fullname())
	fmt.Println(p2.fullname())

	utils.Line()

	s := sum(100, 200, 300)
	fmt.Println("s is", s)
	utils.Line()

	n1 := 10
	n2 := 0
	sum := add(n1, n2)
	d, ok := div(n1, n2)

	fmt.Printf("%d + %d is %d\n", n1, n2, sum)
	if ok {
		fmt.Printf("%d / %d is %d\n", n1, n2, d)
	} else {
		fmt.Printf("Couldn't divide %d by %d\n", n1, n2)
	}
}
