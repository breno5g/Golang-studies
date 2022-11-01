package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	verifyBalance := func(t *testing.T, wallet Wallet, expected Bitcoin) {
		t.Helper()

		result := wallet.Balance()

		if result != expected {
			t.Errorf("result %s, expected %s", result, expected)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		expected := Bitcoin(10)

		verifyBalance(t, wallet, expected)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(150)}

		wallet.Withdraw(Bitcoin(75))

		expected := Bitcoin(75)

		verifyBalance(t, wallet, expected)
	})
}
