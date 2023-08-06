package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to new server!")
	})
	http.HandleFunc("/students", students)

	http.ListenAndServe(":5000", nil)
}

func students(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Student Server!")
}
