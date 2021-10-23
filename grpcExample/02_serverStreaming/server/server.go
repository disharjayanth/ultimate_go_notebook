package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/disharjayanth/ultimate_go_notebook/grpcExample/02_serverStreaming/multiplepb"
	"google.golang.org/grpc"
)

type server struct {
	multiplepb.UnimplementedMultipleServiceServer
}

func (s *server) Multiple(req *multiplepb.MultipleRequest, stream multiplepb.MultipleService_MultipleServer) error {
	num := req.GetNum()

	div := 1

	for {
		if num == int64(div) {
			break
		}
		if num%int64(div) == 0 {
			res := multiplepb.MultipleResponse{
				Num: int64(div),
			}
			div++
			stream.Send(&res)
			time.Sleep(1 * time.Second)
		} else {
			div++
		}
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	fmt.Println("Server running @port:50051")
	if err != nil {
		fmt.Println("Error tcp listener@PORT:50051", err)
		return
	}

	grpcServer := grpc.NewServer()

	multiplepb.RegisterMultipleServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Cannot serve server: %v", err)
		return
	}
}
