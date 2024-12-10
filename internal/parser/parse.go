package parser

import (
	"fmt"
	"strings"

	pb "scaper-demo/proto"

	"github.com/PuerkitoBio/goquery"
)

func parseHTMLContent(htmlContent string) (pb.ParsedPageResponse, error) {
	reader := strings.NewReader(htmlContent)

	// Parse the page
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return pb.ParsedPageResponse{}, fmt.Errorf("failed to parse page: %w", err)
	}

	// Product name from title
	name := doc.Find("div.product_main h1").Text()

	// All from "Product Information" table
	table := doc.Find(".product_page table.table-striped > tbody")

	availability := table.Find("tr:contains(Availability) td").Text()
	upc := table.Find("tr:contains(UPC) td").Text()
	priceExclTax := table.Find("tr:contains(\"Price (excl. tax)\") td").Text()
	tax := table.Find("tr:contains(Tax)").Last().Find("td").Text()

	return pb.ParsedPageResponse{
		Name:         strings.TrimSpace(name),
		Availability: strings.TrimSpace(availability),
		Upc:          strings.TrimSpace(upc),
		PriceExclTax: strings.TrimSpace(priceExclTax),
		Tax:          strings.TrimSpace(tax),
	}, nil
}
