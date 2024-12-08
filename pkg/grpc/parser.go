package grpc

import (
	"context"
	"log"
	pb "scaper-demo/proto"
)

type parserServer struct {
	pb.UnimplementedParserServer
}

func (s *parserServer) ParsePage(ctx context.Context, req *pb.RawPageData) (*pb.ParsedPageResponse, error) {
	log.Print("Received request to parse page")

	return &pb.ParsedPageResponse{
		Name:         "Product 1",
		Availability: "In stock",
		Upc:          "123456789",
		PriceExclTax: "100",
		Tax:          "10",
		Message:      "Success",
	}, nil
}

func NewParserServer() pb.ParserServer {
	return &parserServer{}
}
