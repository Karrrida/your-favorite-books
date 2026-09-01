package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
)

type RegisterBody struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world, you've requested: %s\n", r.URL.Path)
}

func getServerOsInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server use number of cpus: %d\n", runtime.NumCPU())
}

func register(w http.ResponseWriter, r *http.Request) {
	var req RegisterBody
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid params", http.StatusBadRequest)
		return
	}

	fmt.Println(req.Email, req.Name)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/info", getServerOsInfo)
	http.HandleFunc("POST /register", register)

	fmt.Println("Starting server on :8080")

	http.ListenAndServe(":8080", nil)
}
