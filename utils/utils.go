package utils

import (
	"fmt"
	"strconv"
)

// ReadFloatInput reads a float input from the user.
func ReadFloatInput(prompt string) (float64, error) {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)
	return strconv.ParseFloat(input, 64)
}

// ReadStringInput reads a string input from the user.
func ReadStringInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)
	return input, nil
}

// ValidatePIN checks if the provided PIN is valid.
func ValidatePIN(pin string) bool {
	// Simple validation: PIN should be exactly 4 digits
	if len(pin) != 4 {
		return false
	}
	_, err := strconv.Atoi(pin)
	return err == nil
}

// IsValidAccountNumber checks if the provided account number is valid.
func IsValidAccountNumber(accountNumber string) bool {
	// Simple validation: Account number should be exactly 6 digits
	if len(accountNumber) != 6 {
		return false
	}
	_, err := strconv.Atoi(accountNumber)
	return err == nil
}
