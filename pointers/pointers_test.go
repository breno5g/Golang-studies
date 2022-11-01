package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		result := wallet.Balance()

		expected := Bitcoin(10)

		if result != expected {
			t.Errorf("result %s, expected %s", result, expected)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(150)}

		wallet.Withdraw(Bitcoin(75))

		result := wallet.Balance()

		expected := Bitcoin(75)

		if result != expected {
			t.Errorf("result %s, expected %s", result, expected)
		}
	})
}
