package rpc

import (
	"bytes"
	"encoding/json"
	"eth-parser/internal/parser"
	"fmt"
	"net/http"
)

const ethereumRPC = "https://ethereum-rpc.publicnode.com"

// GetBlockNumber retrieves the current block number from the Ethereum blockchain.
func GetBlockNumber() (int, error) {
	// Create the JSON-RPC request body
	requestBody := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_blockNumber", // JSON-RPC method to get the latest block number
		"params":  []interface{}{},   // No parameters required
		"id":      1,
	}

	// Serialize the request body to JSON
	body, _ := json.Marshal(requestBody)

	// Make the HTTP POST request to the Ethereum RPC endpoint
	resp, err := http.Post(ethereumRPC, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	// Decode the JSON-RPC response
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	// Extract the block number from the result (in hexadecimal format)
	hexResult := result["result"].(string)
	var blockNumber int
	fmt.Sscanf(hexResult, "0x%x", &blockNumber) // Convert hex to integer
	return blockNumber, nil
}

// GetBlockTransactions retrieves transactions from a specific block number.
func GetBlockTransactions(blockNumber int) ([]parser.Transaction, error) {
	// Create the JSON-RPC request body
	requestBody := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_getBlockByNumber",                                // JSON-RPC method to get block details
		"params":  []interface{}{fmt.Sprintf("0x%x", blockNumber), true}, // Block number in hexadecimal format and fullTransaction=true
		"id":      1,
	}

	// Serialize the request body to JSON
	body, _ := json.Marshal(requestBody)

	// Make the HTTP POST request to the Ethereum RPC endpoint
	resp, err := http.Post(ethereumRPC, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to make RPC call: %w", err)
	}
	defer resp.Body.Close()

	// Decode the JSON-RPC response
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode RPC response: %w", err)
	}

	// Extract block data from the response
	blockData, ok := result["result"].(map[string]interface{})
	if !ok || blockData == nil {
		return nil, fmt.Errorf("no result found for block number %d", blockNumber)
	}

	// Extract transactions from the block data
	txs, ok := blockData["transactions"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("no transactions found in block number %d", blockNumber)
	}

	// Convert the transactions to the parser.Transaction type
	var transactions []parser.Transaction
	for _, tx := range txs {
		txMap, ok := tx.(map[string]interface{})
		if !ok {
			continue
		}

		transaction := parser.Transaction{
			From:  txMap["from"].(string),                   // Transaction sender address
			To:    txMap["to"].(string),                     // Transaction receiver address
			Value: parser.WeiToEth(txMap["value"].(string)), // Convert Wei to ETH
			Hash:  txMap["hash"].(string),                   // Transaction hash
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
