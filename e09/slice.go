package main

import "fmt"

func main() {
	fmt.Println("Working with Slices")

	var names = []string{"Vinod", "Shyam"}
	var age = 100
	age++

	fmt.Printf("%v %T\n", names, names)
	fmt.Printf("Address of names is %x\n", &names)

	fmt.Printf("%d %d\n", len(names), cap(names))

	names := append(names, "Anu", "Khushi", "Karishma")
	fmt.Printf("%d %d\n", len(names), cap(names))
	fmt.Printf("Address of names is %x\n", &names)

}
