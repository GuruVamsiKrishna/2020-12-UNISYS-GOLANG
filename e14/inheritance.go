package main

import (
	"fmt"
	"go-unisys/utils"
)

type Person struct {
	name  string
	email string
	utils.Address
}

func main() {
	var p1 Person
	p1 = Person{}
	p1.name = "Vinod"
	p1.email = "vinod@vinod.co"
	
	p1.Set("1st crs 1st main", "bangalore", "KA")
	fmt.Println(p1)
	p1.Print()
}
