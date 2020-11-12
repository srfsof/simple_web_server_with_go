package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Browsar")
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server Starting.....")
	http.ListenAndServe(":8000", nil)
}
