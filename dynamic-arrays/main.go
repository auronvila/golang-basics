package main

import "fmt"

func main() {
	prices := []float64{1.1, 1.2, 1.3, 1.4, 1.5}

	// add
	prices = append(prices, 2.0)

	// remove
	prices = append(prices[:1], prices[2:]...)

	fmt.Println(prices)
}
