package tests

import (
	"testing"

	"eth-parser/internal/rpc"
)

func TestGetBlockTransactions(t *testing.T) {
	transactions, err := rpc.GetBlockTransactions(1207)
	if err != nil {
		t.Fatalf("Error fetching block transactions: %v", err)
	}

	if len(transactions) == 0 {
		t.Fatalf("No transactions found in block")
	}

	t.Logf("Transactions: %+v", transactions)
}
