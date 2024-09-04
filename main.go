package main

import (
	"fmt"
	"myatm/account"
	"myatm/auth"
)

func main() {
	accounts := make(map[string]*account.Account)
	accounts["123456"] = &account.Account{AccountNumber: "123456", Pin: 1234, Balance: 999.0}
	accounts["12345678"] = &account.Account{AccountNumber: "12345678", Pin: 4321, Balance: 500.0}

	var accountNumber string
	var pin int64
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
			fmt.Printf("Your balance is: %.2f /-\n", account.CheckBalance())
		case 2:
			var depositAmount float64
			fmt.Println("Enter deposit amount:")
			fmt.Scan(&depositAmount)
			AvilableBalance := account.Deposit(depositAmount)
			fmt.Println("Deposit successful, Total Avilable Balance:", AvilableBalance, "/-")

		case 3:
			var withdrawAmount float64
			fmt.Println("Enter withdrawal amount:")
			fmt.Scan(&withdrawAmount)
			if account.Withdraw(withdrawAmount) {
				fmt.Println("Withdrawal successful Total Avilable Balance:", account.CheckBalance(), "/-")
			} else {
				fmt.Println("Insufficient funds Total Avilable Balance:", account.CheckBalance(), "/-")
			}
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}
