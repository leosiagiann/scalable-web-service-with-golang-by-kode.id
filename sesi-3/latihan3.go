package main

import "fmt"

// this struct is for transaction
type Transaction struct {
	Amount  int
	Total   int
	Product Product
}

// this struct is for product
type Product struct {
	Name  string
	Price int
	Stock int
}

func main() {
	// declare new product with name, price and stock
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

// this function to buy product check stock and return transaction
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
