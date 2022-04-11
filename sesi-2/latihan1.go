package main

import "fmt"

func checkPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	i := 2
	for i <= 50 {
		if checkPrime(i) {
			fmt.Println(i)
		}
		i++
	}
}
