package pointers

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

//Here we are not modifying the value of Wallet so we can use normal reference for Waller
func (w Wallet) Balance() Bitcoin {
	return w.balance
}

//Here we are  modifying the value of Wallet so we  use pointer reference to modify the actual value of Wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance >= amount {
		w.balance -= amount
		return nil
	}
	return errors.New("canot withdraw, insufficient funds")

}
