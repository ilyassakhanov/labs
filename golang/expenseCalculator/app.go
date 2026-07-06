package main

import (
	"errors"
	"fmt"
	"os"
)

const calcFile = "calcFile.txt"

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	// var EBT float64
	// var profit float64
	// var ratio float64

	revenue, err := getUserInput("Revenue Amount: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := getUserInput("Expenses: ")

	if err != nil {
		fmt.Println(err)
		return
	}
	taxRate, err := getUserInput("TaxRate: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	EBT, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	writeCalcToFile(EBT, profit, ratio)

	fmt.Print("EBT: ")
	fmt.Printf("%.1f\n", EBT)

	fmt.Print("Profit: ")
	fmt.Printf("%.1f\n", profit)

	fmt.Print("Ratio: ")
	fmt.Printf("%.3f\n", ratio)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("Value can't be 0 or negative")
	}
	return userInput, nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (EBT float64, profit float64, ratio float64) {
	EBT = revenue - expenses
	profit = EBT - (EBT * taxRate / 100)
	ratio = EBT / profit
	return
}

func writeCalcToFile(EBT, profit, ratio float64) {
	result := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", EBT, profit, ratio)
	os.WriteFile(calcFile, []byte(result), 0644)
}

// Not needed lol
func getBalanceFromFile() (EBT, profit, ratio float64, err error) {
	data, err := os.ReadFile(calcFile)

	if err != nil {
		return 0, 0, 0, errors.New("Failed to find balance file.")
	}

	calcText := string(data)

	n, err := fmt.Sscanf(calcText, "EBT: %f\nprofit: %f\nratio: %f", &EBT, &profit, &ratio)

	if n != 3 || err != nil {
		return 0, 0, 0, errors.New("Failed to parse stored balance value")
	}

	return EBT, profit, ratio, nil
}
