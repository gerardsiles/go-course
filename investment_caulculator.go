package main

import (
	"fmt"
	"math"
)

const InflationRate = 2.5
const ReturnRate = 5.5

func main() {
	var investmentAmount float64 = 1000
	var years float64 = 10

	printPrompt("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	printPrompt("Expected return rate: ")
	fmt.Scan(&returnRate)

	fmt.Print("Investment Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, returnRate, years)
	// futureValue := (investmentAmount) * math.Pow(1 + expectedReturnRate / 100, (years))
	// futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future value (adjusted for inflation): %.2f\n", futureRealValue)

	// fmt.Printf("Future Value: %.2f\nFuture value (adjusted for inflation): %.2f\n", futureValue, futureRealValue)
	// fmt.Println("Future value (adjusted for inflation): ",futureRealValue)
	fmt.Print(formattedFV, formattedFRV)
}

func printPrompt(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, returnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+returnRate/100, (years))
	rfv := fv / math.Pow(1+InflationRate/100, years)
	return fv, rfv
}
