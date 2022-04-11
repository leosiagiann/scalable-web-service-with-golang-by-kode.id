package main

import "fmt"

func main() {
	var product = []map[string]interface{}{
		{
			"name":  "Kalung",
			"stock": 3,
			"price": 10000,
		},
		{
			"name":  "Benang",
			"stock": 3,
			"price": 2000,
		},
	}
	total := 0

	for _, v := range product {
		price := v["price"].(int)
		stock := v["stock"].(int)
		total += price * stock
	}

	fmt.Println(total)
}
