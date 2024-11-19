package handlers

import (
	"encoding/json"
	"net/http"

	"eth-parser/internal/parser"
	"eth-parser/pkg/logging"
)

// TransactionsHandler handles requests to fetch transactions for a subscribed address.
func TransactionsHandler(p *parser.EthereumParser, logger *logging.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		address := r.URL.Query().Get("address")
		if address == "" {
			// Log a warning if the address parameter is missing
			logger.Warn("Address is missing in /transactions request")
			http.Error(w, "Address is required", http.StatusBadRequest)
			return
		}

		// Fetch transactions for the given address
		transactions := p.GetTransactions(address)

		// Check if there are no transactions for the address
		if len(transactions) == 0 {
			logger.Infof("No transactions found for address: %s", address)
			http.Error(w, "No transactions found for this address", http.StatusNotFound)
			return
		}

		// Encode and send transactions as JSON
		if err := json.NewEncoder(w).Encode(transactions); err != nil {
			// Log an error if encoding fails
			logger.Errorf("Failed to encode transactions for address %s: %v", address, err)
			http.Error(w, "Failed to encode transactions", http.StatusInternalServerError)
			return
		}

		// Log successful response
		logger.Infof("Transactions sent for address: %s", address)
	}
}
