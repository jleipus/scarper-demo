package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"scaper-demo/internal/service"
	pb "scaper-demo/proto"

	"google.golang.org/grpc"
)

const baseURL = "http://books.toscrape.com/catalogue/page-%d.html"

func main() {
	service.Run(run)
}

func run(ctx context.Context) error {
	parserAddress := flag.String("parser", "localhost:8080", "The parser service address")
	startSite := flag.String("site", "https://books.toscrape.com/index.html", "The site to scrape")
	flag.Parse()

	conn, err := grpc.NewClient(*parserAddress)
	if err != nil {
		return fmt.Errorf("failed to connect to server: %w", err)
	}

	parserClient := pb.NewParserClient(conn)

	log.Printf("Scraping site %s", *startSite)

	file, err := os.Open(`.\cmd\scraper\example_page.html`)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	response, err := parserClient.ParsePage(ctx, &pb.RawPageData{HtmlContent: string(data)})
	if err != nil {
		return fmt.Errorf("failed to parse page: %w", err)
	}

	log.Printf("Parsed page: %v", response)

	<-ctx.Done()
	return nil
}
