package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// var revenue float64
	var expenses float64
	var taxRate float64

	// fmt.Print("Enter revenue: ")
	// fmt.Scan(&revenue)
	revenue, err := getInput("Enter revenue: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter taxRate: ")
	fmt.Scan(&taxRate)

	// earningsBeforeTax := revenue - expenses
	// profit := earningsBeforeTax * (1 - taxRate / 100)
	// ratio := earningsBeforeTax / profit

	earningsBeforeTax, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("Expenses before tax: %.2f\n", earningsBeforeTax)
	// fmt.Print("Expenses before tax: ")
	// fmt.Println(earningsBeforeTax)
	fmt.Print("Profit: ")
	fmt.Println(profit)
	fmt.Print("Ratio: ")
	fmt.Println(ratio)
	storeResults(earningsBeforeTax, profit, ratio)
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate/100)
	ratio := earningsBeforeTax / profit

	return earningsBeforeTax, profit, ratio
}

func getInput(str string) (float64, error) {
	var val float64
	fmt.Print(str)
	fmt.Scan(&val)

	if val <= 0 {
		return 0, errors.New("value must be positive number")
	}

	return val, nil
}
