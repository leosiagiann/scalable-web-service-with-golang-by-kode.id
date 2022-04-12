package main

import (
	"fmt"
	"latihan/helper"
	"sync"
)

func main() {
	employee := []helper.EmpFunctions{
		helper.NewEmployee("Budi", "IT"),
		helper.NewEmployee("Andi", "HR"),
		helper.NewEmployee("Siti", "Finance"),
		helper.NewEmployee("Karunia", "BE"),
		helper.NewEmployee("Bagas", "FE"),
	}
	var wg sync.WaitGroup
	wg.Add(len(employee))
	for _, emp := range employee {
		go showDataEmployee(emp, &wg)
	}
	wg.Wait()
}

func showDataEmployee(employee helper.EmpFunctions, wg *sync.WaitGroup) {
	fmt.Println(employee.GetName(), "\t", employee.GetDivision())
	wg.Done()
}
