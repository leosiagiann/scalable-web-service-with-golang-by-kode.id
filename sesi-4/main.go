package main

import (
	"fmt"
	"reflect"
)

func main() {
	// circle := helper.NewCircle(5)

	// println(circle)
	dump := []interface{}{
		"leo",
		"dewi",
		2,
	}
	fmt.Println(reflect.ValueOf(dump).Kind())
}
