package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Data struct represents the structure of JSON data expected in the request

type OrderItem struct {
	ItemName                           string  `json:"itemName"`
	ItemDescription                    string  `json:"itemDescription"`
	Price                              float64 `json:"price"`
	UniversalCountryCurrencyIdentifier string  `json:"universalCountryCurrencyIdentifier"`
}

// Define the struct for Order, which contains a slice of OrderItems and a totalTransactionPrice
type Order struct {
	OrderItems            []OrderItem `json:"orderItems"`
	TotalTransactionPrice float64     `json:"totalTransactionPrice"`
}

// Define the struct for the entire transaction
type Transaction struct {
	ObjectID                  string  `json:"objectId"`
	TransactionID             int64   `json:"transactionId"`
	VendorName                string  `json:"vendorName"`
	VendorAddress             string  `json:"vendorAddress"`
	Order                     Order   `json:"order"`
	ReceivedSumFromCustomer   float64 `json:"receivedSumFromCustomer"`
	DeficitToReturnToCustomer float64 `json:"deficitToReturnToCustomer"`
}

// createTransaction simulates creating and returning a Transaction struct
func createTransaction() Transaction {
	return Transaction{
		ObjectID:      "5f8f9b7e8d9e3c4b5a6b7c8d9e0f1a2b3c4d5e6f7a8b9c0d1e2f3a4b5c6d7e8f",
		TransactionID: 1234567890123456,
		VendorName:    "Acme Corporation",
		VendorAddress: "123 Market Street, Springfield, IL, 62701",
		Order: Order{
			OrderItems: []OrderItem{
				{
					ItemName:                           "Wireless Mouse",
					ItemDescription:                    "Bluetooth-enabled mouse with ergonomic design",
					Price:                              29.99,
					UniversalCountryCurrencyIdentifier: "USD",
				},
				{
					ItemName:                           "Keyboard",
					ItemDescription:                    "Mechanical keyboard with RGB lighting",
					Price:                              79.99,
					UniversalCountryCurrencyIdentifier: "USD",
				},
			},
			TotalTransactionPrice: 109.98,
		},
		ReceivedSumFromCustomer:   100.00,
		DeficitToReturnToCustomer: 9.98,
	}
}

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define POST endpoint
	router.POST("/api/transaction", func(c *gin.Context) {
		// Create transaction data
		transaction := createTransaction()
		// Respond with JSON
		c.JSON(http.StatusOK, transaction)
	})

	// GET endpoint at /albums
	router.GET("/albums", getAlbums)
	router.Run("localhost:8080")
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
