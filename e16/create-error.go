package main

import (
	"errors"
	"fmt"
)

type Employee struct {
	Id     int
	Name   string
	Salary float64
}

func (e *Employee) SetId(id int) error {
	if id <= 0 {
		return errors.New("Id must be > 0")
	}
	e.Id = id
	return nil
}

func (e *Employee) SetName(name string) {
	e.Name = name
}

func (e *Employee) SetSalary(salary float64) error {
	if salary < 25000. {
		return errors.New("Salary must be >= 25000")
	}
	e.Salary = salary
	return nil
}

func (e *Employee) Bonus() (float64, error) {
	if e.Salary > 50000 {
		return 0, errors.New("Bonus only for emps with salary <= 50000")
	}

	return e.Salary * 1.5, nil
}

func main() {
	var e1 Employee

	e1.SetId(123)
	e1.SetName("Scott")
	if err := e1.SetSalary(75000.); err != nil {
		fmt.Println("Couldn't set salary:", err)
	}

	fmt.Println(e1)

	if b, err := e1.Bonus(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Bonus is Rs.", b)
	}

}
