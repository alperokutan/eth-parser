package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"eth-parser/internal/handlers"
	"eth-parser/internal/parser"
	"eth-parser/pkg/logging"
)

func TestSubscribeHandler(t *testing.T) {
	logger := logging.NewLogger()
	p := parser.NewEthereumParser(nil)

	// Create the handler
	handler := handlers.SubscribeHandler(p, logger)

	// Simulate a request with a valid address
	req := httptest.NewRequest("GET", "/subscribe?address=0x1234567890abcdef", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	expectedBody := "Subscribed to address: 0x1234567890abcdef\n"
	if w.Body.String() != expectedBody {
		t.Errorf("Expected body %q, got %q", expectedBody, w.Body.String())
	}
}
