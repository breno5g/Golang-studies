package main

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(value int) {
	w.balance += value
}

func (w *Wallet) Balance() int {
	return w.balance
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
