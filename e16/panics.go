package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Need two integers as args")
		return
	}

	n, e := strconv.Atoi(os.Args[1])
	if e != nil {
		fmt.Println("Expecting int")
		return
	}

	d, er := strconv.Atoi(os.Args[2])
	if er != nil {
		fmt.Println("Expecting int")
		return
	}

	q := n / d
	fmt.Printf("%v / %v = %v\n", n, d, q)
}
