package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// fmt.Println("outside the loop, i is", i) // error i is out of scope

	var j int = 5

	for ; j >= 0; j-- {
		fmt.Println("j is", j)
	}

	k := 1

	// while loop of Java/C/C++
	for k <= 5 {
		fmt.Println("k is", k)
		k++
	}

	for a, b := 1, 10; a <= b; {
		fmt.Println("a is", a, "and b is", b)
		a++
		b--
	}

}
