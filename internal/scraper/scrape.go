package scraper

import (
	"context"
	"errors"
	"fmt"
	"io"
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

// Run scrapes the website and handles the product urls.
func Run(ctx context.Context, handler func([]string) error, pageLimit int, delay time.Duration) error {
	page := 1
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		// Get the next page URL
		nextPage := baseURL + fmt.Sprintf(catalogueEndpoint, page)

		// Get the product urls from the page
		urls, err := fetchProductURLs(nextPage)
		if err != nil {
			if err == lastPageErr {
				break
			}

			return fmt.Errorf("failed to scrape page: %w", err)
		}

		// Handle the product urls
		if err := handler(urls); err != nil {
			return fmt.Errorf("failed to handle product urls: %w", err)
		}

		// Wait the delay before the next request
		time.Sleep(delay)

		if pageLimit > 0 && page > pageLimit {
			break
		}
		page++
	}

	return nil
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
func FetchHTMLContent(url string) (string, error) {
	// Fetch the page
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch page: %w", err)
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch page with status: %s", resp.Status)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	// Convert the body to string and return it
	return string(body), nil
}
