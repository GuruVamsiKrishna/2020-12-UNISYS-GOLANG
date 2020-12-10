package main

import "fmt"

func main() {

	var num int
	fmt.Println("num is", num)
	fmt.Print("Enter a number: ")

	fmt.Scanf("%d", &num)

	if num%2 == 0 {
		fmt.Println(num, "is an even number")
	} else {
		fmt.Println(num, "is an odd number")
	}
}
