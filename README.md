# Go Transaction Service

This repository provides a simple web service built in Go using the Gin framework. It includes a single POST endpoint that generates and returns transaction data in JSON format.

## Features

- **Transaction Data API**: A POST endpoint that returns a sample transaction response.
- **Gin Framework**: Built with [Gin](https://github.com/gin-gonic/gin), a fast and lightweight HTTP web framework for Go.

## Prerequisites

- **Go**: Make sure you have [Go installed](https://golang.org/doc/install) (version 1.16+ recommended).
- **Gin**: This project uses Gin as a dependency, which will be installed automatically.

## Getting Started

### 1. Clone the Repository

Clone this repository to your local machine:

```bash
git clone https://github.com/your-username/transaction-service.git
cd transaction-service
```
### 2. Install the required dependencies, including Gin, by running:

```bash
go mod tidy
```
This command will download and install all dependencies specified in go.mod.

### 3. Run the web service, use one of the 2 commands outlined below
```bash
go run main.go
go run . 
```
The server will start on http://localhost:8080.

### 4. Test the Endpoint
You can test the /transaction endpoint using curl or a tool like Postman.
```bash
curl -X POST http://localhost:8080/transaction
```
This command will return a JSON response with the sample transaction data:

```json
{
  "objectId": "5f8f9b7e8d9e3c4b5a6b7c8d9e0f1a2b3c4d5e6f7a8b9c0d1e2f3a4b5c6d7e8f",
  "transactionId": 1234567890123456,
  "vendorName": "Acme Corporation",
  "vendorAddress": "123 Market Street, Springfield, IL, 62701",
  "order": {
    "orderItems": [
      {
        "itemName": "Wireless Mouse",
        "itemDescription": "Bluetooth-enabled mouse with ergonomic design",
        "price": 29.99,
        "universalCountryCurrencyIdentifier": "USD"
      },
      {
        "itemName": "Keyboard",
        "itemDescription": "Mechanical keyboard with RGB lighting",
        "price": 79.99,
        "universalCountryCurrencyIdentifier": "USD"
      }
    ],
    "totalTransactionPrice": 109.98
  },
  "receivedSumFromCustomer": 100.00,
  "deficitToReturnToCustomer": 9.98
}
```
 
