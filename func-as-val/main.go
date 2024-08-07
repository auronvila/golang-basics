package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	// use the function as a parameter and based on the function perform calculations
	transformedNums := transformNumbers(numbers, tripleNumber)

	fmt.Println(transformedNums)
}

func transformNumbers(numbers []int, transformFunc func(int) int) []int {
	var transformedNumbers []int

	for _, val := range numbers {
		transformedNumbers = append(transformedNumbers, transformFunc(val))
	}

	return transformedNumbers
}

func doubleNumber(num int) int {
	return num * 2
}

func tripleNumber(num int) int {
	return num * 3
}
