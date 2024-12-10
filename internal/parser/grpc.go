package parser

import (
	"context"
	"log"
	pb "scaper-demo/proto"
)

type parserServer struct {
	pb.UnimplementedParserServer
}

func NewParserServer() pb.ParserServer {
	return &parserServer{}
}

func (s *parserServer) ParsePage(ctx context.Context, req *pb.RawPageData) (*pb.ParsedPageResponse, error) {
	log.Print("Received request to parse page")

	resp, err := parseHTMLContent(req.HtmlContent)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
