package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/disharjayanth/ultimate_go_notebook/grpcExample/hellopb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	hellopb.UnimplementedHelloServiceServer
}

func (s *server) Hello(ctx context.Context, req *hellopb.Request) (*hellopb.Response, error) {
	reqString := req.GetMessage()
	if reqString == "" {
		return nil, errors.New("request string cannot be empty")
	}

	return &hellopb.Response{
		Message: "This response is from server " + reqString,
	}, nil
}

func main() {
	fmt.Println("Listening @localhost:50051")
	certFile := "grpcExample/ssl/server.crt"
	keyFile := "grpcExample/ssl/server.pem"
	creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
	if sslErr != nil {
		log.Fatalf("Failed to load certificate: %v", sslErr)
		return
	}

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		fmt.Println("Error creating tcp listener at port:50051", err)
		return
	}

	opts := grpc.Creds(creds)
	grpcServer := grpc.NewServer(opts)

	hellopb.RegisterHelloServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
		return
	}
}
