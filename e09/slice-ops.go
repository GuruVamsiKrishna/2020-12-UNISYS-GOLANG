package main

import "fmt"

func deleteAtIndex(data []int, index int) []int {
	return append(data[:index], data[index+1:]...)
}

func main() {
	nums := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Println(nums[3]) // 40
	// fmt.Println(nums[12]) // panic --> index outside range

	// first 5 elements
	fmt.Println(nums[:5])

	// 5th elements onwards (index of 5th element is 4)
	fmt.Println(nums[4:])

	// 4th to 7th element
	fmt.Println(nums[3:7]) // value at 3 is included; value at 7 is excluded

	// nums = append(nums[:3], nums[6:]...)
	nums = deleteAtIndex(nums, 5)
	fmt.Println(nums)

}
