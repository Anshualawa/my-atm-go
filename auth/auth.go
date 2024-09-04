package auth

import "myatm/account"

func Authenticate(accountNumber string, pin int64, accounts map[string]*account.Account) (*account.Account, bool) {
	account, exists := accounts[accountNumber]
	if !exists || account.Pin != pin {
		return nil, false
	}
	return account, true
}
