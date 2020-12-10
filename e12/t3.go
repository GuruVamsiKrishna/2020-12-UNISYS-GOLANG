package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Using channels")

	messages := make(chan string, 50)

	var wg sync.WaitGroup

	wg.Add(3)
	go func(messages chan<- string) {
		fmt.Println("sending hello")
		messages <- "hello"
		fmt.Println("sending hi")
		messages <- "hi"

		wg.Done()
	}(messages)

	go func(messages chan<- string) {
		fmt.Println("sending welcome")
		messages <- "welcome"
		fmt.Println("sending finished")
		messages <- "finished"
		close(messages)
		wg.Done()
	}(messages)

	go func(messages <-chan string) {
		for msg1 := range messages {
			// msg1 := <-messages
			fmt.Println("received message:", msg1)
		}
		wg.Done()
	}(messages)

	wg.Wait()
}
