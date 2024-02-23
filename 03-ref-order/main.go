package main

import (
	"fmt"
)

type order struct {
	quantity int
	price    int
	customer string
	product  string
}

func (o order) printOrder() string {
	total := o.quantity * o.price
	return fmt.Sprintf("Заказ от %s: %dx %s (Итого: $%d)\n", o.customer, o.quantity, o.product, total)
}

func printOrder(o order) string {
	total := o.quantity * o.price
	return fmt.Sprintf("Заказ от %s: %dx %s (Итого: $%d)\n", o.customer, o.quantity, o.product, total)
}

func main() {
	o := order{
		quantity: 12,
		price:    300,
		customer: "Иван",
		product:  "Книга",
	}
	fmt.Println(o.printOrder())
}