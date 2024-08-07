package main

import "fmt"

func main() {
	factNum := factorial(5)
	fmt.Println(factNum)
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}
