package parser

import (
	"fmt"
	"io"
	"strings"

	pb "scaper-demo/proto"

	"github.com/PuerkitoBio/goquery"
)

// ParseHTMLContent reads the HTML content of a page from an io.Reader, parses it, and returns a pb.ParsedPageResponse.
func ParseHTMLContent(reader io.Reader) (pb.ParsedPageResponse, error) {
	// Parse the HTML content of the page
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return pb.ParsedPageResponse{}, fmt.Errorf("failed to parse HTML content: %w", err)
	}

	// Check if the document is empty
	if doc.Find("body").Children().Length() == 0 {
		return pb.ParsedPageResponse{}, fmt.Errorf("no body found in the HTML document")
	}

	// Get product name from title
	name := doc.Find("div.product_main h1").Text()

	// Save "Product Information" table to a variable
	table := doc.Find(".product_page table.table-striped > tbody")

	// Extract fileds from the table
	availability := table.Find("tr:contains(Availability) td").Text()
	upc := table.Find("tr:contains(UPC) td").Text()
	priceExclTax := table.Find("tr:contains(\"Price (excl. tax)\") td").Text()
	tax := table.Find("tr:contains(Tax):not(:contains(Price)) td").Text()

	// Trim leading and trailing white spaces from the extracted fields, and return the result
	return pb.ParsedPageResponse{
		Name:         strings.TrimSpace(name),
		Availability: strings.TrimSpace(availability),
		Upc:          strings.TrimSpace(upc),
		PriceExclTax: strings.TrimSpace(priceExclTax),
		Tax:          strings.TrimSpace(tax),
	}, nil
}
