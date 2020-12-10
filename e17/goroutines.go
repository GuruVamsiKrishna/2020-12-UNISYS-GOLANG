package main

import (
	"fmt"
	"time"
)

func count(name string) {
	fmt.Printf("start of count with %v\n", name)
	for i := 1; true; i++ {
		fmt.Printf("%v   %v\n", name, i)
		time.Sleep(250 * time.Millisecond)
	}
}

func main() {
	fmt.Println("start of main")
	go count("house")
	go count("car")

	// for i := 1; i < 4; i++ {
	// 	fmt.Println("inside main loop")
	// 	time.Sleep(1 * time.Second)
	// }

	fmt.Println("end of main")
	fmt.Scanln()
}
