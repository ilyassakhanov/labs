package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {

	var investmentAmount float64
	var expectedReturnRate = 5.5
	var years float64

	fmt.Print("Years:")
	fmt.Scan(&years)

	fmt.Print("Investment Amount:")
	fmt.Scan(&investmentAmount)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	// var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future value: %.0f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Real Value: %.0f\n", futureRealValue)
	// fmt.Println("Future value: ", futureValue)
	// fmt.Printf("Future value: %.0f\n", futureValue)
	// fmt.Printf("Future Real Value: %.0f\n", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}
func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	// return fv, rfv
	return
}
