package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// func main() {
// 	name := "Vinod"

// 	go func() {
// 		fmt.Println("name is", name)
// 	}()

// 	name = "Shyam"

// 	time.Sleep(1 * time.Second)
// }

func main() {

	mx := sync.Mutex{}

	var words = []string{} // shared resource
	var wg sync.WaitGroup
	s1 := "my name is vinod"
	s2 := "i live in bangalore"
	s3 := "today is the second day of golang training"

	for _, sentence := range []string{s1, s2, s3} {
		wg.Add(1)
		go func(s1 string) {
			for _, word := range strings.Split(s1, " ") {
				mx.Lock()
				words = append(words, word)
				mx.Unlock()
				time.Sleep(10 * time.Millisecond)
			}
			wg.Done()
		}(sentence)
	}

	wg.Wait()
	fmt.Println(len(words))
	fmt.Println(words)
}
