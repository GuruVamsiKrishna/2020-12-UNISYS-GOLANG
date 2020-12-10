package main

import (
	"fmt"
)

// name := "Vinod"
const name string = "Vinod"
const city = "Bangalore"

const (
	state   = "Karnataka"
	country = "India"
)

func main() {
	const name string = "Kumar"
	const pi float64 = 3.14156
	// const sin45 = math.Sin(45)
	fmt.Println("value of pi is", pi)
	fmt.Println("name is", name)
	test()
}

func test() {
	fmt.Println("name is", name)
}
