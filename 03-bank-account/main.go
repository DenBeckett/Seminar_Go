ackage main

import (
	"fmt"
)

type BankAccount struct {
	owner   string
	balance int64
}

func (b *BankAccount) Deposit(amount int64) {
	b.balance += amount
}

func (b *BankAccount) Withdraw(amount int64) bool {
	if b.balance > amount {
		b.balance -= amount
		return true
	}
	return false
}

func (b *BankAccount) Balance() int64 {
	return b.balance
}

func main() {
	account := &BankAccount{
		owner:   "Иван",
		balance: 1000,
	}

	account.Deposit(500)
	fmt.Printf("%s: Баланс - $%d\n", account.owner, account.Balance())

	account.Withdraw(200)
	fmt.Printf("%s: Баланс - $%d\n", account.owner, account.Balance())

	account.Withdraw(1500)
	fmt.Printf("%s: Баланс - $%d\n", account.owner, account.Balance())
}