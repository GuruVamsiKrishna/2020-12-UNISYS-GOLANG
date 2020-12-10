package e20

import (
	"fmt"
	"log"
)

func Factorial(n int) int {
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	return f
}

func main() {
	log.Println("Running main()")
	a, b := 5, Factorial(5)
	fmt.Printf("%v! is %v\n", a, b)
}

func init() {
	log.Println("Starting with init()")
}
