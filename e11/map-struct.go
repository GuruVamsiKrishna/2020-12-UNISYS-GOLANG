package main

import (
	"fmt"
	"go-unisys/utils"
)

type person struct {
	fname, lname, email string
	age                 int
}

func main() {

	p1 := person{"vinod", "kumar", "vinod@vinod.co", 47}
	p2 := person{"harish", "kumar", "harish@xmpl.com", 38}
	p3 := person{"shyam", "sundar", "shyam@xmpl.com", 48}

	// people := make(map[string]person)
	var people = map[string]person{"john@xmpl.com": person{fname: "john", age: 55}}
	people[p1.email] = p1
	people[p2.email] = p2
	people[p3.email] = p3

	fmt.Println("Map of struct types")
	fmt.Println(people)

	shyam, ok := people["shyam@xmpl.com"]
	if ok {
		fmt.Println(shyam)
	}

	utils.Line()

	// use of _ is only to discard the first return value so that we can collect the second one to "val"
	for _, val := range people {
		fmt.Println(val.fname, val.lname)
	}
}
