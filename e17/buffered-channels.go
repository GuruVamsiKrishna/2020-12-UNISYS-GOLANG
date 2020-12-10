package main

import "fmt"

func main() {

	ch := make(chan string, 100)

	ch <- "Hello"
	ch <- "World"
	ch <- "Again"

	// msg1 := <-ch
	// fmt.Println(msg1)

	// msg1 = <-ch
	// fmt.Println(msg1)

	go func() {
		ch <- "Vinod"
		ch <- "Kumar"
		ch <- "Shyam"
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}

	// for {
	// 	msg, open := <-ch
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }
}
