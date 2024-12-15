package scraper_test

import (
	"os"
	"testing"

	"scaper-demo/internal/scraper"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDatabase(t *testing.T) {
	// Create a temporary file for the SQLite database
	tempFile, err := os.CreateTemp("", "testdb_*.db")
	require.NoError(t, err)
	defer os.Remove(tempFile.Name())

	// Initialize the database
	db, err := scraper.NewDatabase(tempFile.Name())
	require.NoError(t, err)
	defer db.Close()

	// Test SaveProducts
	products := []*scraper.Product{
		{Name: "Product1", Availability: "In Stock", Upc: "123456789012", PriceExclTax: "10.00", Tax: "1.00"},
		{Name: "Product2", Availability: "Out of Stock", Upc: "123456789013", PriceExclTax: "20.00", Tax: "2.00"},
	}
	err = db.SaveProducts(products)
	require.NoError(t, err)

	// Test GetProducts with limit
	retrievedProducts, err := db.GetProducts(1)
	require.NoError(t, err)
	assert.Len(t, retrievedProducts, 1)
	assert.Equal(t, "Product1", retrievedProducts[0].Name)

	// Test GetProducts without limit
	retrievedProducts, err = db.GetProducts(0)
	require.NoError(t, err)
	require.Len(t, retrievedProducts, 2)
	assert.Equal(t, "Product1", retrievedProducts[0].Name)
	assert.Equal(t, "Product2", retrievedProducts[1].Name)
}
