package main

import (
	"errors"
	"fmt"
	"latihan/helper"
)

func main() {
	product := helper.Product{
		Name:  "Sapu Lidi",
		Price: 2400,
		Stock: 4,
	}

	var transaction helper.Transaction

	dataChan := make(chan bool, 1)
	defer recovery()
	go buy(&transaction, 3, &product, dataChan)
	if <-dataChan {
		fmt.Println(transaction)
	} else {
		panic(errors.New("error : Stock not enough"))
	}
}

func buy(transaction *helper.Transaction, tot int, product *helper.Product, dataChan chan bool) {
	if product.Stock >= tot {
		transaction.Total = tot * product.Price
		product.Stock -= tot
		transaction.Amount = tot
		transaction.Product = *product
		dataChan <- true
	}
	dataChan <- false
}

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("There is : ", r)
	} else {
		fmt.Println("Transaction Success")
	}
}
