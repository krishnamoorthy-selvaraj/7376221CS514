package main

import (
	"fmt"
	"log"
	"net/http"
)
func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Krishnamoorthy s")
	})
	http.HandleFunc("/rollno", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "7376221CS514")
	})
	http.HandleFunc("/depart", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "computer science and engineering")
	})
	
	fmt.Printf("Server running (port=8080), route: http://localhost:8080/hello\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}