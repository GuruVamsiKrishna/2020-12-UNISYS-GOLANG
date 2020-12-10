package main

import "fmt"

func main() {
	fmt.Println("start")
	defer fmt.Println("end")
	n, d := 123, 22
	a := n / d
	fmt.Println("a is", a)
}
