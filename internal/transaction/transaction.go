package transaction

import "errors"

func PerformDeposit(amount float64) error {
	if amount <= 0 {
		return errors.New("invalid deposit amount")
	}
	return nil
}

func PerformWithdrawal(amount float64) error {
	if amount <= 0 {
		return errors.New("invalid withdrawal amount")
	}
	return nil
}
