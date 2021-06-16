package main

import (
	"log"
	"net/http"
)

const PORT = ":8000"

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	log.Println("Starting server on", PORT)
	log.Fatal(http.ListenAndServe(PORT, mux))
}
