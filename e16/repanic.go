package main

import "fmt"

func divide(n1, n2 int) int {
	fmt.Println("start of divide")

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("RECOVERED:", err)
			panic("Couldn't divide input by zero")
		}
	}()
	result := n1 / n2
	return result
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("RECOVERED:", err)
		}
	}()

	fmt.Println("start of main")
	r := divide(222, 0)
	fmt.Println("r is", r)
	fmt.Println("end of main")
}
