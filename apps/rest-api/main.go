package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from REST API!")
	})
	fmt.Println("REST API listening on :8080")
	http.ListenAndServe(":8080", nil)
}
