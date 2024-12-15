package parser

import (
	"context"
	"io"
	"log"
	pb "scaper-demo/proto"
	"strings"
)

// parserServer is used to implement pb.ParserServer.
type parserServer struct {
	pb.UnimplementedParserServer
	parseFunc parseFunc
}

type parseFunc func(io.Reader) (pb.ParsedPageResponse, error)

// NewParserServer creates a new parser server.
func NewParserServer(parseFunc parseFunc) pb.ParserServer {
	return &parserServer{
		parseFunc: parseFunc,
	}
}

// ParsePage accepts a pb.RawPageData request, parses the provided HTML content, and returns a pb.ParsedPageResponse.
func (s *parserServer) ParsePage(ctx context.Context, req *pb.RawPageData) (*pb.ParsedPageResponse, error) {
	log.Print("Received request to parse page")

	reader := strings.NewReader(req.HtmlContent)

	resp, err := s.parseFunc(reader)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
