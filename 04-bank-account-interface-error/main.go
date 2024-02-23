package main

import (
	"errors"
	"fmt"
	"time"
)

type Operator interface {
	Balance() float64
	Withdraw(amount float64) error
	Deposit(amount float64) error
	Transactions() []Tx
}

var ErrNegativeValue = errors.New("negative value")

var _ error = (*withdrawError)(nil)

type withdrawError struct {
	balance float64
	amount  float64
}

func (b withdrawError) Error() string {
	return fmt.Sprintf("withdraw error, balance: %.2f, withdraw amount: %.2f", b.balance, b.amount)
}

type ActionKind string

const (
	ActionKindIncr ActionKind = "+"
	ActionKindDecr ActionKind = "-"
)

type Tx struct {
	value     float64    // значение на которое изменилось
	action    ActionKind // действие, прибавляем или отнимаем
	createdAt time.Time
}

// Нужно вывести данные транзакции в формате сумма: +-value, time: время создания транзакции
func (t Tx) Print() string {
	return fmt.Sprintf("Сумма: %s%0.1f, time: %s", t.action, t.value, t.createdAt.Format(time.DateTime))
}

var _ Operator = (*txOperator)(nil)

type txOperator struct {
	transactions []Tx
}

func (t *txOperator) Balance() float64 {
	var total float64
	for _, tx := range t.transactions {
		switch tx.action {
		case ActionKindDecr:
			total -= tx.value
		case ActionKindIncr:
			total += tx.value
		default:
		}
	}

	return total
}

func (t *txOperator) Withdraw(amount float64) error {
	balance := t.Balance()
	if balance-amount < 0 {
		return withdrawError{
			balance: balance,
			amount:  amount,
		}
	}
	t.transactions = append(
		t.transactions, Tx{
			value:     amount,
			action:    ActionKindDecr,
			createdAt: time.Now(),
		},
	)
	return nil
}

func (t *txOperator) Deposit(amount float64) error {
	if amount < 0 {
		return ErrNegativeValue
	}

	t.transactions = append(
		t.transactions, Tx{
			value:     amount,
			action:    ActionKindIncr,
			createdAt: time.Now(),
		},
	)

	return nil
}

func (t *txOperator) Transactions() []Tx {
	return t.transactions
}

func main() {
	var op Operator = &txOperator{}
	_ = op
	if err := op.Withdraw(100); err != nil {
		var wErr withdrawError
		if errors.As(err, &wErr) {
			fmt.Printf("Ошибка снятия, текущий баланс: %.2f, попытка снятия %.2f\n", wErr.balance, wErr.amount)
		} else {
			fmt.Println("Неизвестная ошибка", err)
		}
	}
	fmt.Println("Текущий баланс", op.Balance())
	if err := op.Deposit(100); err != nil {
		if errors.Is(err, ErrNegativeValue) {
			fmt.Println("Невозможно зачислить отрицательный баланс")
		} else {
			fmt.Println("Неизвестная ошибка", err)
		}
	}
	fmt.Println("Текущий баланс", op.Balance())
	if err := op.Withdraw(100); err != nil {
		var wErr withdrawError
		if errors.As(err, &wErr) {
			fmt.Printf("Ошибка снятия, текущий баланс: %.2f, попытка снятия %.2f", wErr.balance, wErr.amount)
		} else {
			fmt.Println("Неизвестная ошибка", err)
		}
	}
	fmt.Println("Текущий баланс", op.Balance())
}
