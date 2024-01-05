package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")
var ErrNegativeDeposit = errors.New("cannot deposit negative amount")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
func (w *Wallet) Deposit(amount Bitcoin) error {
	if amount < 0 {
		return ErrNegativeDeposit
	}
	w.balance += amount
	return nil

}
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func main() {
	w := Wallet{}
	w2 := Wallet{}
	w.balance = 0
	w.Deposit(10)
	w2.Deposit(20)
	w2.Withdraw(100)
	fmt.Println(w.Balance())
	fmt.Println(w2.Balance())
}
