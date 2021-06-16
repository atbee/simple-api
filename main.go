package main

import (
	"log"
	"net/http"
)

const PORT = ":8000"

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func helloSomeone(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new ..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello-world", helloWorld)
	mux.HandleFunc("/hello-someone", helloSomeone)

	log.Println("Starting server on", PORT)
	log.Fatal(http.ListenAndServe(PORT, mux))
}
