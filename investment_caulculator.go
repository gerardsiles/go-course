package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64 = 1000
	expectedReturnRate := 5.5
	var years float64 = 10

	outputText("Investment Amount: ")
	// & is a pointer to the variable
	fmt.Scan(&investmentAmount)

	outputText("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Investment Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount,expectedReturnRate, years)
	// futureValue := (investmentAmount) * math.Pow(1 + expectedReturnRate / 100, (years))
	// futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future value (adjusted for inflation): %.2f\n", futureRealValue)
	
	// fmt.Printf("Future Value: %.2f\nFuture value (adjusted for inflation): %.2f\n", futureValue, futureRealValue)
	// fmt.Println("Future value (adjusted for inflation): ",futureRealValue)
	fmt.Print(formattedFV, formattedFRV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1 + expectedReturnRate / 100, (years))
	rfv := fv / math.Pow(1 + inflationRate / 100, years)
	return fv, rfv
}