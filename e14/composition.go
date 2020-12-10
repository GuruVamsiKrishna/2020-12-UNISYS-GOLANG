package main

import "fmt"

type Address struct {
	street string
	city   string
	state  string
}

type Person struct {
	name    string
	email   string
	address Address
}

func main() {
	var p1 Person
	p1 = Person{}
	p1.name = "Vinod"
	p1.email = "vinod@vinod.co"
	p1.address.city = "Bangalore"
	p1.address.street = "1st crs 1st main"
	p1.address.state = "Karnataka"
	fmt.Println(p1)
}
