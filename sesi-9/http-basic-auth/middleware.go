package main

import "net/http"

const USERNAME = "leonardo"
const PASSWORD = "leonardo"

func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()

	if !ok {
		w.Write([]byte("something went wrong\n"))
		return false
	}

	isValid := (username == USERNAME && password == PASSWORD)
	if !isValid {
		w.Write([]byte("invalid username or password\n"))
	}

	return true
}

func AllowOnlyGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		w.Write([]byte("only GET method is allowed\n"))
		return false
	}

	return true
}
