package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"

	"example.com/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println("err")
		fmt.Println("---------")
		panic("Can't continue, sorry")
	}

	fmt.Println("Welcome to Bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	// for i := 0; i < 2; i++
	for {
		presentOptions()

		var choice int

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		fmt.Println("You chose:", choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)

		case 2:
			var depositAmount float64
			fmt.Println("Deposit amount: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid deposit amount. Must be greater than 0")
				continue

			} else {
				accountBalance += depositAmount
				fmt.Println("Balance updated! New amount:", accountBalance)
				fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			}

		case 3:
			var withdrawalAmount float64
			fmt.Println("Withdrawal amount: ")
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid withdrawal amount. Must be greater than 0")
				continue
			} else if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. Insufficient funds")
				continue
			} else {
				accountBalance -= withdrawalAmount
				fmt.Println("Balance updated! New amount:", accountBalance)
				fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			}
		default:
			fmt.Println("Goodbye! ")
			return
			// break - can't use inside switch statement,
		}
	}
}
