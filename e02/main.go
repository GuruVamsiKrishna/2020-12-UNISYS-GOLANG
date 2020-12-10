package main

import (
	"fmt"
	"os"
)

func main() {
	// var name string
	// name = "Vinod"

	// var name string = "Vinod"

	// var name = "Vinod" // type inference

	name := "Vinod Kumar" // type inference

	fmt.Printf("Hello, %v\n", name)
	fmt.Printf("%v %T\n", name, name)

	if len(os.Args) > 1 {
		fmt.Println("This is the command line arguments: ", os.Args)
	}
}
