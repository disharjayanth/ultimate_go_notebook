package main

import (
	"context"
	"fmt"
	"io"

	"github.com/disharjayanth/ultimate_go_notebook/grpcExample/04_biStreaming/maxpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("error dailing @port:50051", err)
		return
	}

	client := maxpb.NewMaxServiceClient(conn)

	maxOfInts := []int{10, 2, 12, 4, 15, 2, 4, 5, 6, 100}

	clientStream, err := client.Max(context.Background())
	if err != nil {
		fmt.Println("client stream error", err)
		return
	}

	waitC := make(chan struct{})

	go func() {
		for _, v := range maxOfInts {
			req := maxpb.MaxRequest{
				Num: int64(v),
			}
			if err = clientStream.Send(&req); err != nil {
				fmt.Println("Error sending request to server:", err)
				return
			}
		}

		if err = clientStream.CloseSend(); err != nil {
			fmt.Println("erorr closing client stream", err)
			return
		}
	}()

	go func() {
		for {
			resp, err := clientStream.Recv()
			if err == io.EOF {
				// send of serve stream
				close(waitC)
				return
			}
			if err != nil {
				fmt.Println("error receving response from server:", err)
				close(waitC)
				return
			}

			fmt.Println(resp.GetRes())
		}
	}()

	<-waitC
}
