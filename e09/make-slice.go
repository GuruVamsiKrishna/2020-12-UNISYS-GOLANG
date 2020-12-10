package main

import "fmt"

func main() {
	names := make([]string, 0, 100)
	fmt.Printf("len = %d, cap = %d\n", len(names), cap(names))
	names = append(names, "Vinod")
	fmt.Printf("len = %d, cap = %d\n", len(names), cap(names))
	names = append(names, "Shyam", "Harish")
	fmt.Printf("len = %d, cap = %d\n", len(names), cap(names))
	names = append(names, "Kishore")
	fmt.Printf("len = %d, cap = %d\n", len(names), cap(names))
	names = append(names, "John", "Jane", "Kane")
	fmt.Printf("len = %d, cap = %d\n", len(names), cap(names))
}
