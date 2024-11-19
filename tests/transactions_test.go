package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"eth-parser/internal/handlers"
	"eth-parser/internal/parser"
	"eth-parser/pkg/logging"
)

func TestTransactionsHandler(t *testing.T) {
	logger := logging.NewLogger()
	p := parser.NewEthereumParser(nil)

	// Subscribe an address and add mock transactions
	address := "0x1234567890abcdef"
	p.Subscribe(address)
	p.GetTransactions(address)

	mockTransaction := parser.Transaction{
		From:  "0xabcdef",
		To:    address,
		Value: "1.5",
		Hash:  "0xdeadbeef",
	}
	p.GetTransactions(address)

	// Create the handler
	handler := handlers.TransactionsHandler(p, logger)

	// Simulate a request for the address
	req := httptest.NewRequest("GET", "/transactions?address="+address, nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	// Validate the response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var transactions []parser.Transaction
	if err := json.NewDecoder(w.Body).Decode(&transactions); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if len(transactions) != 1 {
		t.Errorf("Expected 1 transaction, got %d", len(transactions))
	}

	if transactions[0].Hash != mockTransaction.Hash {
		t.Errorf("Expected transaction hash %s, got %s", mockTransaction.Hash, transactions[0].Hash)
	}
}
