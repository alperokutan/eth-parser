package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	notify := func(address string, tx Transaction) {
		fmt.Printf("Notification: New transaction for %s: %+v\n", address, tx)
	}

	parser := NewEthereumParser(notify)

	go parser.detectNewBlocks()

	http.HandleFunc("/subscribe", func(w http.ResponseWriter, r *http.Request) {
		address := r.URL.Query().Get("address")
		if address == "" {
			http.Error(w, "Address is required", http.StatusBadRequest)
			return
		}

		subscribed := parser.Subscribe(address)
		if subscribed {
			fmt.Fprintf(w, "Subscribed to address: %s\n", address)
		} else {
			fmt.Fprintf(w, "Already subscribed to address: %s\n", address)
		}
	})

	http.HandleFunc("/transactions", func(w http.ResponseWriter, r *http.Request) {
		address := r.URL.Query().Get("address")
		if address == "" {
			http.Error(w, "Address is required", http.StatusBadRequest)
			return
		}

		transactions := parser.GetTransactions(address)
		if err := json.NewEncoder(w).Encode(transactions); err != nil {
			http.Error(w, "Failed to encode transactions", http.StatusInternalServerError)
		}
	})

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
