package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
)

const ethereumRPC = "https://ethereum-rpc.publicnode.com"

func getBlockNumber() (int, error) {
	requestBody := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_blockNumber",
		"params":  []interface{}{},
		"id":      1,
	}

	body, _ := json.Marshal(requestBody)
	resp, err := http.Post(ethereumRPC, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	hexResult := result["result"].(string)
	var blockNumber int
	fmt.Sscanf(hexResult, "0x%x", &blockNumber)
	return blockNumber, nil
}

func getBlockTransactions(blockNumber int) ([]Transaction, error) {
	requestBody := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_getBlockByNumber",
		"params":  []interface{}{fmt.Sprintf("0x%x", blockNumber), true},
		"id":      1,
	}

	body, _ := json.Marshal(requestBody)
	resp, err := http.Post(ethereumRPC, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	blockData, ok := result["result"].(map[string]interface{})
	if !ok || blockData == nil {
		return nil, fmt.Errorf("failed to get block data")
	}

	txs, ok := blockData["transactions"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("failed to parse transactions")
	}

	var transactions []Transaction
	for _, tx := range txs {
		txMap, ok := tx.(map[string]interface{})
		if !ok {
			continue
		}

		transaction := Transaction{
			From:  txMap["from"].(string),
			To:    txMap["to"].(string),
			Value: weiToEth(txMap["value"].(string)),
			Hash:  txMap["hash"].(string),
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

// Converts Wei value to ETH
func weiToEth(wei string) string {
	weiBig, _ := new(big.Int).SetString(wei, 10)
	ethValue := new(big.Float).Quo(new(big.Float).SetInt(weiBig), big.NewFloat(1e18))
	return ethValue.Text('f', 18)
}
