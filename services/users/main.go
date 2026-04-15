package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from users service!")
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server running on :8081")
	http.ListenAndServe(":8081", nil)
}