package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"latihan/models"
)

var users = []models.User{
	{
		Name: "John",
		Age:  30,
	},
	{
		Name: "Jane",
		Age:  25,
	},
}

func main() {
	port := ":9999"

	http.HandleFunc("/api/users", getUsers)
	http.HandleFunc("/api/users/add", addUsers)
	http.HandleFunc("/api/users/name", getUserByName)

	http.HandleFunc("/users", listUsersHTML)

	fmt.Println("server running at port", port)

	http.ListenAndServe(port, nil)
}

func getUsers(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	if r.Method == http.MethodGet {
		json.NewEncoder(rw).Encode(users)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(rw).Encode(map[string]interface{}{
		"status": http.StatusMethodNotAllowed,
		"msg":    "method not allowed",
	})
}

func addUsers(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	if r.Method == http.MethodPost {
		var req models.User

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(rw).Encode(map[string]interface{}{
				"status": http.StatusBadRequest,
				"msg":    err.Error(),
			})
			return
		}

		users = append(users, req)

		json.NewEncoder(rw).Encode(map[string]interface{}{
			"status": http.StatusOK,
			"msg":    "create user success",
		})
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(rw).Encode(map[string]interface{}{
		"status": http.StatusMethodNotAllowed,
		"msg":    "method not allowed",
	})
}

func getUserByName(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	if r.Method == http.MethodGet {
		query := r.URL.Query().Get("name")

		var user *models.User
		for _, u := range users {
			if u.Name == query {
				user = &u
				break
			}
		}

		if user == nil {
			rw.WriteHeader(http.StatusNotFound)
			json.NewEncoder(rw).Encode(map[string]interface{}{
				"status": http.StatusNotFound,
				"msg":    "user with name - " + query + " not found",
			})
			return
		}

		json.NewEncoder(rw).Encode(map[string]interface{}{
			"status": http.StatusOK,
			"msg":    user,
		})
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(rw).Encode(map[string]interface{}{
		"status": http.StatusMethodNotAllowed,
		"msg":    "method not allowed",
	})
}

type DataHTML struct {
	Title string
	Data  interface{}
}

func listUsersHTML(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t, err := template.ParseFiles("views/user.html")
		if err != nil {
			fmt.Println(err)
			return
		}

		t.Execute(rw, DataHTML{
			Title: "Users",
			Data:  users,
		})
		return
	}

	http.Error(rw, "Method not allowed", http.StatusBadRequest)
}
