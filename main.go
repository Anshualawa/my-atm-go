package main

import (
	"fmt"
	"myatm/account"
	"myatm/auth"
	"myatm/operations"
	"myatm/utils"
)

func main() {
	accounts := make(map[string]*account.Account)
	accounts["123456"] = &account.Account{AccountNumber: "123456", Pin: "1234", Balance: 1000.0}

	accountNumber, _ := utils.ReadStringInput("Enter account number: ")
	pin, _ := utils.ReadStringInput("Enter PIN: ")

	if !utils.IsValidAccountNumber(accountNumber) || !utils.ValidatePIN(pin) {
		fmt.Println("Invalid account number or PIN format")
		return
	}

	acc, authenticated := auth.Authenticate(accountNumber, pin, accounts)
	if !authenticated {
		fmt.Println("Authentication failed")
		return
	}

	var option int
	for {
		fmt.Println("\n1. Check Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")
		fmt.Scan(&option)

		switch option {
		case 1:
			fmt.Printf("Your balance is: %.2f\n", operations.CheckBalance(acc))
		case 2:
			depositAmount, _ := utils.ReadFloatInput("Enter deposit amount: ")
			operations.Deposit(acc, depositAmount)
			fmt.Println("Deposit successful")
		case 3:
			withdrawAmount, _ := utils.ReadFloatInput("Enter withdrawal amount: ")
			if operations.Withdraw(acc, withdrawAmount) {
				fmt.Println("Withdrawal successful")
			} else {
				fmt.Println("Insufficient funds")
			}
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}
