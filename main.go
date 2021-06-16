package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = ":8000"

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", hello)

	log.Fatal(http.ListenAndServe(PORT, nil))
}
