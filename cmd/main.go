package main

import (
	"net/http"

	"eth-parser/internal/handlers"
	"eth-parser/internal/parser"
	"eth-parser/pkg/logging"
)

func main() {
	logger := logging.NewLogger()

	// Create an EthereumParser and attach a notification logger
	notify := func(address string, tx parser.Transaction) {
		logger.Infof("New transaction detected for %s: %+v", address, tx)
	}

	ethParser := parser.NewEthereumParser(notify)

	// Start a goroutine to detect new blocks
	go func() {
		defer func() {
			// Recover in case of panic and log the error
			if r := recover(); r != nil {
				logger.Errorf("Recovering from panic: %v", r)
			}
		}()
		ethParser.DetectNewBlocks()
	}()

	// Register HTTP handlers
	http.HandleFunc("/subscribe", handlers.SubscribeHandler(ethParser, logger))
	http.HandleFunc("/transactions", handlers.TransactionsHandler(ethParser, logger))

	// Start the HTTP server
	logger.Info("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
