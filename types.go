package main

type Transaction struct {
	From  string
	To    string
	Value string // transaction amount
	Hash  string // transaction hash
}

type Parser interface {
	GetCurrentBlock() int
	Subscribe(address string) bool
	GetTransactions(address string) []Transaction
}
