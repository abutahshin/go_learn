package bnk

import (
	"errors"
	"fmt"
	"time"
)

type AuditInfo struct {
	createdAt    time.Time
	lastModified time.Time
}
type AccountNumber string

type BankAccount struct {
	accountNumber AccountNumber
	balance       float64
	AuditInfo     // Embedding AuditInfo struct
}

func (ba BankAccount) DisplayBalance() {
	fmt.Printf("Account Number: %s, Balance: €%.2f\n", ba.accountNumber, ba.balance)
}

func (ba *BankAccount) Deposit(amount float64) {
	ba.balance += amount
}

func (ba *BankAccount) Withdraw(amount float64) error {
	if ba.balance < amount {
		return errors.New("insufficient funds")
	}
	ba.balance -= amount
	return nil
}

func NewBankAccount(num AccountNumber) *BankAccount {
	return &BankAccount{
		accountNumber: num,
		AuditInfo:     AuditInfo{createdAt: time.Now(), lastModified: time.Now()},
	}
}
