package account

import (
	"time"
)

type Customer struct {
	ID   int64
	Name string
	City string
}

type Transaction struct {
	ID     int64
	Type   string
	Amount float64
	Date   time.Time
}

type Account struct {
	ID           int64
	Balance      float64
	Customer     Customer
	Transactions []Transaction
}

func (a *Account) GetBalance() float64 {
	return a.Balance
}

func (a *Account) GetTransactions() []Transaction {
	return a.Transactions
}

func (a *Account) GenerateMonthlyStatement() []Transaction {
	return a.Transactions
}

type SavingsAccount struct {
	Account
}

type CheckingAccount struct {
	Account
}
