package handlers

import (
	"fmt"
	"net/http"

	"eth-parser/internal/parser"
	"eth-parser/pkg/logging"
)

// SubscribeHandler handles requests to subscribe an address.
func SubscribeHandler(p *parser.EthereumParser, logger *logging.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		address := r.URL.Query().Get("address")
		if address == "" {
			// Log a warning if the address parameter is missing
			logger.Warn("Address is missing in /subscribe request")
			http.Error(w, "Address is required", http.StatusBadRequest)
			return
		}

		// Attempt to subscribe the address
		subscribed := p.Subscribe(address)
		if subscribed {
			// Log if a new address is subscribed
			logger.Infof("Address subscribed: %s", address)
			fmt.Fprintf(w, "Subscribed to address: %s\n", address)
		} else {
			// Log if the address is already subscribed
			logger.Infof("Address already subscribed: %s", address)
			fmt.Fprintf(w, "Already subscribed to address: %s\n", address)
		}
	}
}
