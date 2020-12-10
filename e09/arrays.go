package main

import "fmt"

func main() {
	fmt.Println("Working arrays")

	var nums [3]int
	nums[0] = 12
	nums[1] = 22
	nums[2] = 39

	fmt.Println("nums is", nums)
	fmt.Println("len(nums) is", len(nums))

	const size = 5
	var names [size]string
	fmt.Println("len(names) is", len(names))
	fmt.Println("names is", names)

	var cities = [...]string{"bangalore", "mysore", "shivamogga", "mandya"}
	fmt.Printf("%v %T\n", cities, cities)

	fmt.Println("Cities are...:")
	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

	for index, value := range cities {
		fmt.Printf("%d --> %v\n", index, value)
	}

	// _ is called 'write-only' variable
	for _, n := range nums {
		fmt.Println(n)
	}

	name := "vinod"
	for _, c := range name {
		fmt.Println(c, string(c))
	}

}
