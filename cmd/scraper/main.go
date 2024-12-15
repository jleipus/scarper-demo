package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"scaper-demo/internal/scraper"
	"scaper-demo/internal/service"
	pb "scaper-demo/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const baseURL = "http://books.toscrape.com/catalogue/page-%d.html"

func main() {
	service.Run(run)
}

func run(ctx context.Context) error {
	parserAddress := flag.String("parser", "localhost:8080", "Parser service address")
	pageLimit := flag.Int("limit", 0, "Page limit")
	scrapingDelay := flag.Duration("delay", time.Second, "Delay between requests")
	pageRetryLimit := flag.Int("retry", 3, "Number of retries for each page")
	sqliteDatabase := flag.String("db", "products.db", "SQLite database file")
	flag.Parse()

	// Create new database connection
	db, err := scraper.NewDatabase(*sqliteDatabase)
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}
	defer db.Close()

	// Create new gRPC connection
	conn, err := grpc.NewClient(*parserAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return fmt.Errorf("failed to connect to server: %w", err)
	}

	// Create new parser client
	parserClient := pb.NewParserClient(conn)

	// Run the scraper
	scraper.Run(ctx,
		scraper.FetchHTMLContent, // readProductPage
		func(content []byte) (scraper.Product, error) {
			// Send the HTML content to the parser service, get the parsed page response
			resp, err := parserClient.ParsePage(ctx, &pb.RawPageData{HtmlContent: string(content)})

			// Return the parsed product, converted to the internal representation
			return scraper.TransformProductRPC(resp), err
		}, // parseProductPage
		func(products []*scraper.Product) {
			err := db.SaveProducts(products)
			if err != nil {
				log.Printf("failed to save products: %v", err)
			}
		}, // saveProducts
		func(err error) {
			log.Printf("error: %v", err)
		}, // logError
		*pageLimit,
		*pageRetryLimit,
		*scrapingDelay,
	)

	// Display the first 20 products
	products, err := db.GetProducts(20)
	if err != nil {
		return fmt.Errorf("failed to get products: %w", err)
	}

	displayProducts(products)

	<-ctx.Done()
	return nil
}

func displayProducts(products []*scraper.Product) {
	truncateString := func(s string, length int) string {
		if len(s) > length {
			return s[:length-3] + "..."
		}

		return s
	}

	nameLength := 30
	availabilityLength := 23
	upcLength := 16
	priceLength := 8
	taxLength := 8

	rowFormat := fmt.Sprintf("%%-%ds | %%-%ds | %%-%ds | %%-%ds | %%-%ds\n", nameLength, availabilityLength, upcLength, priceLength, taxLength)
	fmt.Printf(rowFormat, "Name", "Availability", "UPC", "Price", "Tax")
	for _, product := range products {
		name := truncateString(product.Name, nameLength)
		availability := truncateString(product.Availability, availabilityLength)
		upc := truncateString(product.Upc, upcLength)
		price := truncateString(product.PriceExclTax, priceLength)
		tax := truncateString(product.Tax, taxLength)
		fmt.Printf(rowFormat, name, availability, upc, price, tax)
	}
}
