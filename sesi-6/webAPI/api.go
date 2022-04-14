package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
}

var employees = []Employee{
	{ID: 1, Name: "John", Age: 30, Division: "IT"},
	{ID: 2, Name: "Jane", Age: 25, Division: "HR"},
}

var port = ":8080"

func main() {
	http.HandleFunc("/employees", getEmployees)
	http.HandleFunc("/employee", createEmployee)

	fmt.Println("server running at port", port)
	http.ListenAndServe(port, nil)
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		template, err := template.ParseFiles("templates/employees.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		template.Execute(w, employees)
		return
	}

	http.Error(w, "Method not allowed", http.StatusBadRequest)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		age := r.FormValue("age")
		division := r.FormValue("division")

		intAge, err := strconv.Atoi(age)

		if err != nil {
			http.Error(w, "Invalid age", http.StatusBadRequest)
			return
		}

		newEmployee := Employee{
			ID:       len(employees) + 1,
			Name:     name,
			Age:      intAge,
			Division: division,
		}

		employees = append(employees, newEmployee)

		json.NewEncoder(w).Encode(newEmployee)
		return
	}

	http.Error(w, "Method not allowed", http.StatusBadRequest)
}
