package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount, years, expectedReturnRate float64

	fmt.Print("Please enter an investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Please enter years: ")
	fmt.Scan(&years)
	fmt.Print("Please enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureVal := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealVal := futureVal / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureVal)
	fmt.Println(futureRealVal)
}
