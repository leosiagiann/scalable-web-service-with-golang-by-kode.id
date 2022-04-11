package main

import "fmt"

type Transaction struct {
	Amount  int
	Total   int
	Product Product
}

type Product struct {
	Name  string
	Price int
	Stock int
}

func main() {
	product := Product{
		Name:  "Sapu Lidi",
		Price: 2400,
		Stock: 4,
	}

	var transaction Transaction

	if buy(&transaction, 2, &product) {
		fmt.Println(transaction)
	} else {
		fmt.Println("Stock tidak mencukupi")
	}
}

func buy(transaction *Transaction, tot int, product *Product) bool {
	if product.Stock >= tot {
		transaction.Total = tot * product.Price
		product.Stock -= tot
		transaction.Amount = tot
		transaction.Product = *product
		return true
	}
	return false
}
