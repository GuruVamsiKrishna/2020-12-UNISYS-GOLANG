package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var wg2 sync.WaitGroup

	fmt.Println("start of main")
	wg.Add(1)
	wg2.Add(2)

	go func() {
		count("house", 9)
		wg.Done()
	}()

	go func() {
		count("car", 3)
		wg2.Done()
	}()

	go func() {
		count("elephant", 3)
		wg2.Done()
	}()

	wg.Wait()
	wg2.Wait()

	fmt.Println("end of main")
}

func count(name string, count int) {
	fmt.Printf("start of count with %v\n", name)
	for i := 1; i <= count; i++ {
		fmt.Printf("%v   %v\n", name, i)
		time.Sleep(150 * time.Millisecond)
	}
}
