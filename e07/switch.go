package main

import "fmt"

func main() {
	var month int

	fmt.Print("Enter value for a month (1-12): ")
	fmt.Scanf("%d", &month)

	switch month {
	case 2:
		fmt.Println("28 or 29 days")
	case 4, 6, 9, 11:
		fmt.Println("30 days")
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("31 days")
	default:
		fmt.Println("Invalid value for month - ", month)
	}
}
