package main

import (
	"errors"
	"fmt"
	"time"
)

type AuditInfo struct {
	CreditAt     time.Time
	LastModified time.Time
}
type bankaccount struct {
	AccountNumber string
	Balance       float64
	AuditInfo
}

func (ba bankaccount) Display() {
	fmt.Println("Account Number:", ba.AccountNumber, "Current Balance:", ba.Balance)
	fmt.Println("Last Credited At:", ba.CreditAt, " And Last Modified At:", ba.LastModified)
}
func (ba *bankaccount) Deposit(amount float64) {
	ba.Balance += amount
	ba.CreditAt = time.Now()
	ba.LastModified = time.Now()
}
func (ba *bankaccount) Withdraw(amount float64) error {
	if amount < ba.Balance {
		return errors.New("Insufficient funds")
	}
	ba.Balance -= amount
	ba.LastModified = time.Now()
	return nil
}
func main() {
	account := bankaccount{AccountNumber: "123456789",
		Balance: 0.0, AuditInfo: AuditInfo{CreditAt: time.Now(), LastModified: time.Now()},
	}
	fmt.Printf("Before Deposit Amount: ")
	account.Display()
	account.Deposit(1000.0)
	fmt.Printf("After Deposit Amount: ")
	account.Display()

}
