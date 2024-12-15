package scraper

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	baseURL           = "https://books.toscrape.com/catalogue/"
	catalogueEndpoint = "page-%d.html"
)

// lastPageErr is returned when the last page is reached.
//
// Alternatively, a custom error type could be implemented, however for this use case,
// a simple error variable is sufficient.
var lastPageErr = errors.New("last page")

// readProductPageFunc is a function that takes a URL and reads returns the HTML conent of the page.
type readProductPageFunc func(string) ([]byte, error)

// parseProductPageFunc is a function that parses the HTML conent of a product page and returns a Product struct.
type parseProductPageFunc func([]byte) (Product, error)

// saveProductsFunc is a function that concurrently saves a slice of Product objects to the database.
type saveProductsFunc func([]*Product)

type logErrorFunc func(error)

// Run scrapes the website and handles the product urls.
// TODO: Implement retry logic
func Run(ctx context.Context, readProductPage readProductPageFunc, parseProductPage parseProductPageFunc,
	saveProducts saveProductsFunc, logError logErrorFunc, pageLimit, retryLimit int, delay time.Duration) {
	page := 1
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		// Get the next page URL
		nextPage := baseURL + fmt.Sprintf(catalogueEndpoint, page)

		log.Printf("Scraping page %d: %s", page, nextPage)

		// Get the product urls from the page
		urls, err := fetchProductURLs(nextPage)
		if err != nil {
			if err == lastPageErr {
				break
			}

			logError(fmt.Errorf("failed to scrape page %d: %w", page, err))
		} else {
			log.Printf("Found %d products", len(urls))

			products := make([]*Product, 0, len(urls))
			for _, url := range urls {
				// Get the HTML content for a product page
				html, err := readProductPage(url)
				if err != nil {
					logError(fmt.Errorf("failed to fetch page: %w", err))
					continue
				}

				// Parse the page
				response, err := parseProductPage([]byte(html))
				if err != nil {
					logError(fmt.Errorf("failed to parse page: %w", err))
					continue
				}

				products = append(products, &response)
			}

			// Save the products to the database
			go saveProducts(products)
		}

		// Wait the delay before the next request
		time.Sleep(delay)

		if pageLimit > 0 && page > pageLimit {
			break
		}
		page++
	}

	return
}

// fetchProductURLs makes a request to the URL and extracts the product links from the HTML content.
func fetchProductURLs(url string) ([]string, error) {
	// Fetch the page
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch page: %w", err)
	}
	defer resp.Body.Close()

	// Stop if we encounter a 404
	if resp.StatusCode == http.StatusNotFound {
		return nil, lastPageErr
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch page with status: %s", resp.Status)
	}

	return extractProductURLs(resp.Body)
}

func extractProductURLs(reader io.Reader) ([]string, error) {
	// Parse the page
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to parse page: %w", err)
	}

	var urls []string
	doc.Find("ol.row li article.product_pod h3 a").Each(func(index int, item *goquery.Selection) {
		href, exists := item.Attr("href")
		if exists {
			urls = append(urls, baseURL+href)
		}
	})

	return urls, nil
}

// FetchHTMLContent fetches the HTML content from the provided links and calls the handler function.
func FetchHTMLContent(url string) ([]byte, error) {
	// Fetch the page
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch page: %w", err)
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch page with status: %s", resp.Status)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Convert the body to string and return it
	return body, nil
}
