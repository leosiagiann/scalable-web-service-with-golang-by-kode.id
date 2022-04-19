package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FullName string   `json:"full_name"`
	Email    string   `json:"email"`
	Age      int      `json:"age"`
	Division Division `json:"division"`
}

type Division struct {
	Name string `json:"name"`
}

func main() {
	var jsonString = `
		[
			{
				"full_name": "Leonardo Siagian",
				"email": "leonardosiagian2001@gmail.com",
				"age" : 20,
				"division": {
					"name": "Software Engineering"
				}
			},
			{
				"full_name": "Putri Sitompul",
				"email": "putisitompul718@gmail.com",
				"age" : 21,
				"division": {
					"name": "Quality Assurance"
				}
			}
		]
	`

	var result []Employee

	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(result)
}
