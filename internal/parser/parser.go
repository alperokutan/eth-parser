package parser

import (
	"sync"
	"time"

	"eth-parser/internal/rpc"
)

type EthereumParser struct {
	currentBlock int
	addresses    map[string][]Transaction
	mu           sync.Mutex
	onNewTx      func(address string, tx Transaction)
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

func (p *EthereumParser) DetectNewBlocks() {
	for {
		latestBlock, err := rpc.GetBlockNumber()
		if err != nil {
			// Error handling için log eklenebilir
			time.Sleep(10 * time.Second)
			continue
		}

		if latestBlock > p.currentBlock {
			for b := p.currentBlock + 1; b <= latestBlock; b++ {
				p.processBlock(b)
			}
			p.currentBlock = latestBlock
		}

		time.Sleep(10 * time.Second)
	}
}

func (p *EthereumParser) processBlock(blockNumber int) {
	transactions, err := rpc.GetBlockTransactions(blockNumber)
	if err != nil {
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
