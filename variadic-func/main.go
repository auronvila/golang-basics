package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	//summed := sumup(1, 2, 3, 4, 5)
	summed := sumup(numbers...)

	fmt.Println(summed)
}

func sumup(nums ...int) int {
	res := 0

	for _, num := range nums {
		res += num
	}
	return res
}
