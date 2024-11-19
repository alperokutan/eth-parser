package parser

import (
	"math/big"
)

// Transaction represents a transaction on the Ethereum blockchain.
type Transaction struct {
	From  string `json:"from"`  // Address initiating the transaction
	To    string `json:"to"`    // Target address of the transaction
	Value string `json:"value"` // Transaction amount (stored in ETH)
	Hash  string `json:"hash"`  // Transaction hash
}

// WeiToEth converts a value in Wei (smallest unit of Ether) to ETH.
func WeiToEth(wei string) string {
	weiBig, ok := new(big.Int).SetString(wei, 10)
	if !ok {
		return "0" // Return 0 ETH if the input is invalid
	}
	ethValue := new(big.Float).Quo(new(big.Float).SetInt(weiBig), big.NewFloat(1e18))
	return ethValue.Text('f', 18) // Return the ETH value with 18 decimal places
}
