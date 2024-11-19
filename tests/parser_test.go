package tests

import (
	"testing"

	"eth-parser/internal/parser"
)

func TestSubscribe(t *testing.T) {
	p := parser.NewEthereumParser(nil)
	address := "0x1234567890abcdef"
	if !p.Subscribe(address) {
		t.Errorf("Subscribe failed for address: %s", address)
	}
	if !p.Subscribe(address) {
		t.Errorf("Address should already be subscribed: %s", address)
	}
}
