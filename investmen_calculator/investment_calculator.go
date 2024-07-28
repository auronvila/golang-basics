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

	// using sprintf to return the string that we are formating and using it to print to the terminal.
	formattedOutput := fmt.Sprintf("Future value: %.3f \nFuture Val adjusted for inflation: %.3f", futureVal, futureRealVal)

	fmt.Printf(formattedOutput)
}
