package wallet

import (
  "errors"
  "fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin float64

type Stringer interface {
  String() string
}

func (b Bitcoin) String() string {
  return fmt.Sprintf("%.2f BTC", b)
}

type Wallet struct { 
  balance Bitcoin
}

func (w *Wallet) Balance() Bitcoin {
  return w.balance
}

func (w *Wallet) Deposit(amount Bitcoin){
  w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
  if amount > w.balance {
    return ErrInsufficientFunds
  }
  w.balance -= amount
  return nil
}
