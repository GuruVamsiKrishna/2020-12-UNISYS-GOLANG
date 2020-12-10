package main

import "fmt"

func main() {

	marks := 18.23

	switch {
	case marks < 35.:
		fmt.Println("Fail")
	case marks < 50.:
		fmt.Println("Passing grade")
	case marks < 60:
		fmt.Println("Second class")
	default:
		fmt.Println("First class / distinction")
	}
}
