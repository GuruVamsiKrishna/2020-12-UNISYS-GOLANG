package main

import "fmt"

func factorial(n int, ch chan<- int) {
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}

	ch <- f
	close(ch)
}

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go factorial(5, ch1)
	go factorial(8, ch2)

	for {
		select {
		case f1 := <-ch1:
			fmt.Println(f1)
		case f2 := <-ch2:
			fmt.Println(f2)
		}
	}
}
