package main

import "fmt"

const (
	sunday = iota
	monday
	tuesday
	wednessday
	thursday
	friday
	saturday
)

const (
	jan = 1 + iota
	feb
	mar
	apr
	may
	jun
)

const (
	isAdmin      = 1 << iota // 1 * 2^0    			0000 0001
	isSupervisor             // 1 * 2^1				0000 0010
	isManager                // 1 * 2^2				0000 0100
	isDirector               // 1*2^3				0000 1000
)

func main() {
	fmt.Printf("sunday = %v\n", sunday)
	fmt.Printf("monday = %v\n", monday)
	fmt.Printf("tuesday = %v\n", tuesday)
	fmt.Println(jan, feb, mar, apr)

	vinodRoles := isAdmin | isManager // 0000 0101
	fmt.Printf("vinodRoles is %b\n", vinodRoles)

	if vinodRoles&isDirector == 1 {
		// 0101 <- vinodRoles
		// 1000 <- isDirector
		// 0000 <- vinodRoles & isDirector
		fmt.Println("Vinod is a director")
	}
	if vinodRoles&isAdmin == 1 {
		// 0101
		// 0001
		// 0001
		fmt.Println("Vinod is an admin")
	}
}
