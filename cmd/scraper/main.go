package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
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
	parserAddress := flag.String("parser", "localhost:8080", "The parser service address")
	pageLimit := flag.Int("limit", 0, "The page limit")
	scrapingDelay := flag.Duration("delay", time.Second, "The delay between requests")
	flag.Parse()

	conn, err := grpc.NewClient(*parserAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect to server: %w", err)
	}

	parserClient := pb.NewParserClient(conn)

	err = scraper.Run(ctx, func(urls []string) error {
		for _, url := range urls {
			html, err := scraper.FetchHTMLContent(url)
			if err != nil {
				return fmt.Errorf("failed to fetch page: %w", err)
			}

			response, err := parserClient.ParsePage(ctx, &pb.RawPageData{HtmlContent: html})
			if err != nil {
				return fmt.Errorf("failed to parse page: %w", err)
			}

			bytes, err := json.MarshalIndent(response, "", "  ")
			if err != nil {
				return fmt.Errorf("failed to marshal response: %w", err)
			}

			fmt.Println(string(bytes))
		}

		return nil
	}, *pageLimit, *scrapingDelay)
	if err != nil {
		return fmt.Errorf("failed to run scraper: %w", err)
	}

	return nil
}
