package helper

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
