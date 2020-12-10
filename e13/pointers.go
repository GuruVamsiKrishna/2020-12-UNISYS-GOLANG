package main

import "fmt"

func incr(n *int) {
	(*n)++
	fmt.Printf("incr >> n=%d addr=%x\n", *n, n)
}

func main() {

	var n int = 1
	fmt.Printf("main >> n=%d addr=%x\n", n, &n)

	incr(&n)
	incr(&n)
	incr(&n)

	fmt.Printf("main >> n=%d addr=%x\n", n, &n)

}
