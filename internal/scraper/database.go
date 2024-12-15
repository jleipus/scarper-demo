package scraper

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Product struct {
	gorm.Model

	Name         string
	Availability string
	Upc          string
	PriceExclTax string
	Tax          string
}

// ProductDatabase is an interface for a product database.
type ProductDatabase interface {
	// SaveProducts saves a product to the database.
	SaveProducts(products []*Product) error

	// GetProducts returns requested number of products from the database.
	// If the number of products is 0, all products are returned.
	GetProducts(int) ([]*Product, error)

	// Close closes the database connection.
	Close() error
}

type database struct {
	db *gorm.DB
}

func NewDatabase(connectionString string) (ProductDatabase, error) {
	db, err := gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	return database{db: db}, nil
}

func (d database) SaveProducts(products []*Product) error {
	tx := d.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}}, // Key column
		DoUpdates: clause.AssignmentColumns([]string{
			"name", "availability", "upc", "price_excl_tax", "tax",
		}), // Columns needed to be updated
	}).Create(&products)

	return tx.Error
}

func (d database) GetProducts(limit int) ([]*Product, error) {
	var products []*Product
	tx := d.db
	if limit > 0 {
		tx = tx.Limit(limit)
	}

	err := tx.Find(&products).Error
	return products, err
}

func (d database) Close() error {
	sqlDB, err := d.db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
