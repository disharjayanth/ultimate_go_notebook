package main

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/disharjayanth/ultimate_go_notebook/grpcExample/01_basicClientServer/hellopb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	// Certificate Authority Trust Certificate
	certFile, _ := filepath.Abs("ssl/ca.crt")
	creds, err := credentials.NewClientTLSFromFile(certFile, "")
	if err != nil {
		fmt.Println("Error while loading CA trust certificate:", err)
		return
	}

	opts := grpc.WithTransportCredentials(creds)
	cc, err := grpc.Dial("localhost:50051", opts)
	defer cc.Close()

	if err != nil {
		fmt.Println("Cannot dail: %v", err)
		return
	}

	client := hellopb.NewHelloServiceClient(cc)

	req := &hellopb.Request{
		Message: "Hello dishar",
	}

	res, err := client.Hello(context.Background(), req)
	if err != nil {
		fmt.Println("Error getting response:", err)
		return
	}

	fmt.Println("Response:", res.GetMessage())
}
