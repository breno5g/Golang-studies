package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(value Bitcoin) {
	w.balance += value
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(value Bitcoin) error {
	if value > w.balance {
		return errors.New("não é possível retirar: saldo insuficiente")
	}

	w.balance -= value
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// func main() {
// 	var creature string = "shark"
// 	// Without the * and the &, the value is just passed as a value
// 	// Using them, it passes a reference to the location in memory
// 	var pointer *string = &creature

// 	creature = "tutu"

// 	fmt.Println("creature =", creature)
// 	fmt.Println("pointer =", pointer)
// 	fmt.Println("pointer =", *pointer)
// }
