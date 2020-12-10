package main

import "fmt"

func month() int {
	var m int
	fmt.Print("Enter a month: ")
	fmt.Scanf("%d", &m)
	return m
}

func main() {

	if m := month(); m == 2 {
		fmt.Println("28 or 29 days")
	} else if m == 4 || m == 6 || m == 9 || m == 11 {
		fmt.Println("30 days")
	} else {
		fmt.Println("31 days")
	}

	fmt.Println("value of m is", m)
}
