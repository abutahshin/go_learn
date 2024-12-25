package bnk

import (
	"fmt"
	"time"
)

type Customer struct {
	name     string
	accounts []BankAccount
	AuditInfo
}

func (c *Customer) AddAccount(account BankAccount) {
	c.accounts = append(c.accounts, account)
}

// Method to display all accounts of a customer
func (c Customer) DisplayAccounts() {
	fmt.Printf("Customer: %s\n", c.name)
	for _, account := range c.accounts {
		account.DisplayBalance()
	}
}

func NewCustomer(name string) *Customer {
	return &Customer{
		name:      name,
		AuditInfo: AuditInfo{createdAt: time.Now(), lastModified: time.Now()},
	}
}
