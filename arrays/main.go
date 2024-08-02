package main

import "fmt"

func main() {
	myMap := map[string]int{
		"test":  1,
		"test2": 2,
	}
	prices := [4]float64{11, 2.3, 32.1, 13}

	priceSlice := prices[0:2]

	fmt.Println(prices)
	fmt.Println(priceSlice)
	fmt.Println(myMap)
}
