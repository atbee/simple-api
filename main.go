package main

import (
	"encoding/json"
	"fmt"
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
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var b map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var name interface{}
	if name = b["name"]; name == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	message := fmt.Sprintf("Hello, %v!", name)
	w.Write([]byte(message))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello-world", helloWorld)
	mux.HandleFunc("/hello-someone", helloSomeone)

	log.Println("Starting server on", PORT)
	log.Fatal(http.ListenAndServe(PORT, mux))
}
