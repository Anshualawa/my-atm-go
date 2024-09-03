package operations

import "myatm/account"

// CheckBalance checks the balance of the given account.
func CheckBalance(acc *account.Account) float64 {
	return acc.CheckBalance()
}

// Deposit deposits the specified amount into the given account.
func Deposit(acc *account.Account, amount float64) {
	acc.Deposit(amount)
}

// Withdraw withdraws the specified amount from the given account.
func Withdraw(acc *account.Account, amount float64) bool {
	return acc.Withdraw(amount)
}
