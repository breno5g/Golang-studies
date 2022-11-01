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

	verifyError := func(t *testing.T, err error, expected string) {
		t.Helper()
		if err == nil {
			t.Fatal("esperava um erro, mas nenhum ocorreu")
		}

		resultado := err.Error()

		if resultado != expected {
			t.Errorf("resultado %s, expected %s", resultado, expected)
		}
	}

	verifyUnexistedError := func(t *testing.T, err error) {
		t.Helper()
		if err != nil {
			t.Fatal("esperava um erro, mas nenhum ocorreu")
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		expected := Bitcoin(10)

		verifyBalance(t, wallet, expected)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		result := wallet.Withdraw(Bitcoin(10))

		expected := Bitcoin(10)

		verifyBalance(t, wallet, expected)
		verifyUnexistedError(t, result)
	})

	t.Run("withdraw money with insufficient balance", func(t *testing.T) {
		initialBalance := Bitcoin(20)
		wallet := Wallet{balance: initialBalance}

		err := wallet.Withdraw(Bitcoin(75))

		verifyError(t, err, "não é possível retirar: saldo insuficiente")
	})
}
