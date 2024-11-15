package main

import (
	"fmt"
	"sync"
	"time"
)

type EthereumParser struct {
	currentBlock int
	addresses    map[string][]Transaction
	mu           sync.Mutex
	onNewTx      func(address string, tx Transaction) // Notification callback
}

func NewEthereumParser(callback func(address string, tx Transaction)) *EthereumParser {
	return &EthereumParser{
		currentBlock: 0,
		addresses:    make(map[string][]Transaction),
		onNewTx:      callback,
	}
}

func (p *EthereumParser) GetCurrentBlock() int {
	return p.currentBlock
}

func (p *EthereumParser) Subscribe(address string) bool {
	p.mu.Lock()
	defer p.mu.Unlock()
	if _, exists := p.addresses[address]; !exists {
		p.addresses[address] = []Transaction{}
		return true
	}
	return false
}

func (p *EthereumParser) GetTransactions(address string) []Transaction {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.addresses[address]
}

func (p *EthereumParser) detectNewBlocks() {
	for {
		latestBlock, err := getBlockNumber()
		if err != nil {
			fmt.Println("Error fetching block number:", err)
			time.Sleep(10 * time.Second)
			continue
		}

		if latestBlock > p.currentBlock {
			fmt.Printf("New block detected: %d\n", latestBlock)
			for b := p.currentBlock + 1; b <= latestBlock; b++ {
				p.processBlock(b)
			}
			p.currentBlock = latestBlock
		}

		time.Sleep(10 * time.Second)
	}
}

func (p *EthereumParser) processBlock(blockNumber int) {
	// Fetch and process transactions in the new block
	transactions, err := getBlockTransactions(blockNumber)
	if err != nil {
		fmt.Println("Error fetching transactions for block:", blockNumber, err)
		return
	}

	for _, tx := range transactions {
		p.mu.Lock()
		if _, exists := p.addresses[tx.To]; exists {
			p.addresses[tx.To] = append(p.addresses[tx.To], tx)
			if p.onNewTx != nil {
				p.onNewTx(tx.To, tx)
			}
		}
		p.mu.Unlock()
	}
}
