package account

type Account struct {
	AccountNumber string
	Pin           int64
	Balance       float64
}

func (a *Account) CheckBalance() float64 {
	return a.Balance
}

func (a *Account) Deposit(amount float64) float64 {
	a.Balance += amount
	return a.Balance
}

func (a *Account) Withdraw(amount float64) bool {
	if amount > a.Balance {
		return false
	}
	a.Balance -= amount
	return true
}
