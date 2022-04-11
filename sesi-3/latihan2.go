package main

import "fmt"

func main() {
	isPrime := func(num int) bool {
		for i := 2; i < num; i++ {
			if num%i == 0 {
				return false
			}
		}
		return true
	}

	nums := []int{10, 13, 11, 5, 20, 27, 30, 35, 40, 45, 50, 97}
	totalPrime := test(nums, isPrime)
	fmt.Println(totalPrime)
}

func test(nums []int, isPrime func(num int) bool) (result []int) {
	for _, num := range nums {
		if isPrime(num) {
			result = append(result, num)
		}
	}
	return
}
