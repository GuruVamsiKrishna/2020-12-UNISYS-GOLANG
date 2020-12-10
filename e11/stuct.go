package main

import "fmt"

type Employee struct {
	id     int
	name   string
	salary float64
}

// this is method and not a function
func (e *Employee) print() {
	fmt.Printf("ID       : %v\n", e.id)
	fmt.Printf("Name     : %v\n", e.name)
	fmt.Printf("Salary   : %v\n", e.salary)
	fmt.Println()
}

func main() {
	var e1 Employee
	e1 = Employee{id: 1, salary: 55000., name: "Ramesh"}
	e2 := Employee{2, "Harish", 64000.}
	e3 := Employee{}
	e3.id = 33
	e3.name = "Sujay"
	e3.salary = 22000
	e1.print()
	e2.print()
	e3.print()

}
