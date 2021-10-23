package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/disharjayanth/ultimate_go_notebook/grpcExample/03_clientStreaming/avgpb"
	"google.golang.org/grpc"
)

type server struct {
	avgpb.UnimplementedAvgServiceServer
}

func (s *server) Avg(stream avgpb.AvgService_AvgServer) error {
	var sum int64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			res := avgpb.AvgResponse{
				Result: float32(sum / 2),
			}
			if err = stream.SendAndClose(&res); err != nil {
				fmt.Println("Error sending response to client:", err)
				return err
			}
			break
		}
		if err != nil {
			fmt.Println("Error receving client stream", err)
			return err
		}
		num := req.GetNum()
		sum += num
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		fmt.Println("Error creating tcp listener", err)
		return
	}

	grpcServer := grpc.NewServer()

	avgpb.RegisterAvgServiceServer(grpcServer, &server{})

	fmt.Println("Server listening @port:50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error server serving @port:50051 %v\n", err)
	}
}
