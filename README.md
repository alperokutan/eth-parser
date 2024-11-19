# Ethereum Parser

A robust and scalable Ethereum blockchain parser to track transactions for subscribed addresses. It uses the Ethereum JSON-RPC interface to fetch block and transaction data.

## Features

- Track incoming/outgoing transactions for specific Ethereum addresses.
- Expose HTTP endpoints for:
  - Subscribing to Ethereum addresses.
  - Retrieving transactions for subscribed addresses.
- Modular and extensible design.
- Uses Zap for structured logging.
- Converts transaction values from Wei to ETH.
- Includes automated tests for core functionalities.

---

## Installation
```bash
git clone https://github.com/your-repo/ethereum-parser.git
cd ethereum-parser
go mod tidy
go run cmd/main.go