package main

import "fmt"

func factorial(n int, ch chan int) {
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}

	ch <- f
}

func main() {
	ch := make(chan int)
	go factorial(5, ch)

	f := <-ch
	fmt.Println("factorial of 5 is", f)
}
