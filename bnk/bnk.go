package bnk

import "fmt"

type Bank struct {
	name      string
	customers []Customer
}

// Method to add a new customer to a bank
func (b *Bank) AddCustomer(customer Customer) {
	b.customers = append(b.customers, customer)
}

// Method to display all customers and their accounts in a bank
func (b Bank) DisplayCustomers() {
	fmt.Printf("Bank: %s\n", b.name)
	for _, customer := range b.customers {
		customer.DisplayAccounts()
	}
}

func NewBank(name string) *Bank {
	return &Bank{name: "GoBank"}
}
