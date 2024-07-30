package main

import (
	"fmt"
)

func main() {
	accountBalance := 1000.0
	fmt.Println("Welcome to Go bank")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println(accountBalance)
		} else if choice == 2 {
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount, the amount has to be greater than 0")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Your updated balance: ", accountBalance)
		} else if choice == 3 {
			fmt.Print("Your withdraw amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount < 0 {
				fmt.Println("Invalid amount, the amount has to be greater than 0")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount, you cannot withdraw more than you have.")
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("Your updated balance: ", accountBalance)
		} else if choice == 4 {
			fmt.Println("You exited the banking operation successfully")
			break
		} else {
			fmt.Println("You entered an invalid input and got exited by the operation")
			break
		}
	}

	fmt.Println("Thank you for using our bank")
}
