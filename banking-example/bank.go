package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
		fmt.Println("----------")
		// panic("Can't execute program")
	}

	var choice int

	fmt.Println("Welcome to Go Banking!")
	fmt.Println("Reach us at", randomdata.PhoneNumber())
	for {
		presentOptions()

		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")

			var depositAmount float64
			fmt.Scan(&depositAmount)

			accountBalance += depositAmount

			fmt.Println("Balance updated! New amount: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("Withdrawal amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Not enough funds in this account.")
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated. New amount: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

		default:
			fmt.Println("Exiting operation")
			fmt.Println("Thanks for choosing our bank.")

			return
		}

		// if choice == 1 {
		// 	fmt.Println("Your balance is: ", accountBalance)
		// } else if choice == 2 {
		// 	fmt.Print("Your deposit: ")

		// 	var depositAmount float64
		// 	fmt.Scan(&depositAmount)

		// 	accountBalance += depositAmount

		// 	fmt.Println("Balance updated! New amount: ", accountBalance)
		// } else if choice == 3 {
		// 	fmt.Print("Withdrawal amount: ")
		// 	var withdrawalAmount float64
		// 	fmt.Scan(&withdrawalAmount)

		// 	if withdrawalAmount <= 0 {
		// 		fmt.Println("Invalid amount. Must be greater than 0.")
		// 		continue
		// 	}

		// 	if withdrawalAmount > accountBalance {
		// 		fmt.Println("Not enough funds in this account.")
		// 		continue
		// 	}

		// 	accountBalance -= withdrawalAmount
		// 	fmt.Println("Balance updated. New amount: ", accountBalance)
		// } else {
		// 	fmt.Println("Exiting operation")
		// 	// return
		// 	break
		// }
	}
	// fmt.Println("Thanks for choosing our bank.")

}
