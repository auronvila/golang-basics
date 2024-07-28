package main

import "fmt"

func main() {
	revenue, expenses, taxRate := 0.0, 0.0, 0.0

	fmt.Print("enter your revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("enter your expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("enter your taxRate: ")
	fmt.Scan(&taxRate)
	ebt := revenue - expenses
	taxes := ebt * taxRate
	eat := ebt - taxes

	fmt.Println("Earning before taxes: ", ebt, "dollars Earning after taxes", eat)
}
