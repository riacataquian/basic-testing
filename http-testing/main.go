package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// curl http://localhost:8080/ping
func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

// curl http://localhost:8080/hello?person=john
func helloHandler(w http.ResponseWriter, r *http.Request) {
	person := r.FormValue("person")

	if person == "" {
		http.Error(w, "missing required parameter/s", http.StatusBadRequest)
		return
	}

	io.WriteString(w, fmt.Sprintf("Hello, %s!", person))
}
