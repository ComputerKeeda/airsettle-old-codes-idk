package main

import (
	handlers "github.com/airchains-network/settlementfaucet/internal/handlers"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/faucet", handlers.FaucetHandler)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
