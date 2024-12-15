package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "scaper-demo/proto"

	"scaper-demo/internal/parser"
	"scaper-demo/internal/service"

	"google.golang.org/grpc"
)

func main() {
	service.Run(run)
}

func run(ctx context.Context) error {
	serverIP := flag.String("ip", "localhost", "IP address of the server")
	serverPort := flag.Int("port", 8080, "Port of the server")
	flag.Parse()

	serverAddress := fmt.Sprintf("%s:%d", *serverIP, *serverPort)
	listener, err := net.Listen("tcp", serverAddress)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	go func() {
		<-ctx.Done()
		log.Println("Shutting down parser service")
		grpcServer.GracefulStop()
	}()

	pb.RegisterParserServer(grpcServer, parser.NewParserServer(parser.ParseHTMLContent))

	log.Printf("Parser service running on %s", serverAddress)
	if err := grpcServer.Serve(listener); err != nil {
		return err
	}

	<-ctx.Done()
	return nil
}
