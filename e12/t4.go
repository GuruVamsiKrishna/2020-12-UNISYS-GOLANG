package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan int, 10)
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		for n := range ch {
			f := 1
			for i := 2; i <= n; i++ {
				f *= i
			}
			fmt.Printf("factorial of %d is %d\n", n, f)
		}
		wg.Done()
	}()

	go func() {
		nums := []int{5, 6, 7, 8}
		for _, n := range nums {
			ch <- n
		}
		close(ch)
		wg.Done()
	}()

	wg.Wait()
}
