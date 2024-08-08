package customer

import "time"

type Transaction struct {
	ID     int64
	Type   string
	Amount float64
	Date   time.Time
}

type Customer struct {
	ID   int64
	Name string
	City string
}

func (c *Customer) GetTransactionReport(month time.Month, year int) []Transaction {
	return nil
}
