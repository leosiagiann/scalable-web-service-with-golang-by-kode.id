package main

import "fmt"

func main() {
	province := []string{"Jakarta", "Bali", "Yogyakarta", "Surabaya"}
	greeting := make(chan string, len(province))

	go func(greeting chan<- string) {
		for _, p := range province {
			greeting <- p
		}
	}(greeting)

	fmt.Println(greeting)
	fmt.Println(greeting)
	fmt.Println(greeting)

	close(greeting)
}
