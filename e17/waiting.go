package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("start of main")
	wg.Add(1)
	go count("house", 5) // new Thread(r->{}).start()

	wg.Add(1)
	go count("car", 7) // new Thread(r->{}).start()

	wg.Wait() // t1.join(); t2.join()
	fmt.Println("end of main")
}

func count(name string, count int) {
	defer func() {
		fmt.Printf("Done counting %v\n", name)
		wg.Done()
	}()

	fmt.Printf("start of count with %v\n", name)
	for i := 1; i <= count; i++ {
		fmt.Printf("%v   %v\n", name, i)
		time.Sleep(150 * time.Millisecond)
	}
}
