package utils

import "fmt"

func HelloAgain() {
	fmt.Println("Hello, again!")
}

func Line() {
	fmt.Println("--------------------------------------")
}

type Address struct {
	street string
	city   string
	state  string
}

func (a *Address) Set(street, city, state string) {
	a.street = street
	a.city = city
	a.state = state
}

func (a *Address) Print() {
	fmt.Println(a.street)
	fmt.Println(a.city)
	fmt.Println(a.state)
}
