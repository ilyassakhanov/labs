package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile("balance.txt")

	if err != nil {
		return 1000, errors.New("Failed to find balance file.")
	}

	balaceText := string(data)
	balance, err := strconv.ParseFloat(balaceText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value")
	}
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceStr := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceStr), 0644)
}

func main() {
	var accountBalance, err = getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println("err")
		fmt.Println("---------")
		panic("Can't continue, sorry")
	}

	fmt.Println("Welcome to Bank!")

	// for i := 0; i < 2; i++
	for {
		fmt.Println("Please choose operation:")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

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
				writeBalanceToFile(accountBalance)
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
				writeBalanceToFile(accountBalance)
			}
		default:
			fmt.Println("Goodbye! ")
			return
			// break - can't use inside switch statement,
		}
	}
}
