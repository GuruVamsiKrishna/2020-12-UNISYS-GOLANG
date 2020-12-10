package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int8 // max --> 2^7-1  min --> - (2^8)
	n = -128

	n2 := 200

	f1 := 3.

	// a := 2
	// b := 2.
	// fmt.Println("a+b is", a+b)

	fmt.Printf("%v %T\n", n, n)
	fmt.Printf("%v %T\n", f1, f1)
	fmt.Println("n + n2 is ", int(n)+n2)

	name := "Vinod"
	age := 47

	msg := name + " is " + strconv.Itoa(age) + " years old."
	fmt.Println(msg)

}
