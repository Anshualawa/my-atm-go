package main

import (
	"fmt"
	"myatm/account"
	"myatm/auth"
)

func main() {
	accounts := make(map[string]*account.Account)
	accounts["123456"] = &account.Account{AccountNumber: "123456", Pin: "1234", Balance: 1000.0}

	var accountNumber, pin string
	fmt.Println("Enter account number:")
	fmt.Scan(&accountNumber)
	fmt.Println("Enter PIN:")
	fmt.Scan(&pin)

	account, authenticated := auth.Authenticate(accountNumber, pin, accounts)
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
			fmt.Printf("Your balance is: %.2f\n", account.CheckBalance())
		case 2:
			var depositAmount float64
			fmt.Println("Enter deposit amount:")
			fmt.Scan(&depositAmount)
			account.Deposit(depositAmount)
			fmt.Println("Deposit successful")
		case 3:
			var withdrawAmount float64
			fmt.Println("Enter withdrawal amount:")
			fmt.Scan(&withdrawAmount)
			if account.Withdraw(withdrawAmount) {
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
