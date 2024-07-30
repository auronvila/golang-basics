package main

import "fmt"

func main() {
	revenue, expenses, taxRate := 0.0, 0.0, 0.0

	fmt.Print("enter your revenue: ")
	fmt.Scan(&revenue)
	if revenue <= 0 {
		panic("revenue cannot be less than or equal to 0")
	}
	fmt.Print("enter your expenses: ")
	fmt.Scan(&expenses)
	if expenses <= 0 {
		panic("expenses cannot be less than or equal to 0")
	}
	fmt.Print("enter your taxRate: ")
	fmt.Scan(&taxRate)
	if taxRate <= 0 {
		panic("taxRate cannot be less than or equal to 0")
	}
	ebt := revenue - expenses
	taxes := ebt * taxRate
	eat := ebt - taxes

	fmt.Println("Earning before taxes: ", ebt, "dollars Earning after taxes", eat)
}
