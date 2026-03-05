package main

import (
	"errors"
	"fmt"
)

type account struct {
	balance float64
}

type chahat struct {
	chugh bool
}

func (a account) Error() string {
	return fmt.Sprintf("Invalid Balance: %v", a.balance)
}

type bankingOperation interface {
	withdrawal(amount float64) (float64, error)
	credit(amount float64) (float64, error)
}

func (a account) withdrawal(amount float64) (float64, error){
	if a.balance < amount {
		return float64(0),errors.New("insufficient balance")
	}
	a.balance -= amount
	return a.balance, nil
}

func (a account) credit(amount float64) (float64, error){
	if amount < 0 {
		return 0, a;
	}
	a.balance += amount
	return a.balance, nil
}


func performBankingOperation(b bankingOperation,amount float64, operation string) (float64, error){
	// panic("Banking operation interrupted !!!");
	// log.Fatal("Banking operation interrupted !!!");
	switch operation {
	case "withdrawal":
		return b.withdrawal(amount)
	case "credit":
		return b.credit(amount)
	default:
		return 0.0, nil
	}
}

