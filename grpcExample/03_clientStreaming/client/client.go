package main

import (
	"context"
	"fmt"

	"github.com/disharjayanth/ultimate_go_notebook/grpcExample/03_clientStreaming/avgpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Error dailing server @port:50051", err)
	}

	client := avgpb.NewAvgServiceClient(conn)

	sliceOfInts := []int{2, 4, 6, 8}

	clientStream, err := client.Avg(context.Background())
	if err != nil {
		fmt.Println("client stream error", err)
		return
	}

	for _, v := range sliceOfInts {
		req := avgpb.AvgRequest{
			Num: int64(v),
		}

		if err = clientStream.Send(&req); err != nil {
			fmt.Println("Error sending client stream request:", err)
			return
		}
	}

	res, err := clientStream.CloseAndRecv()
	if err != nil {
		fmt.Println("Error receving server response:", err)
		return
	}

	fmt.Println(res.GetResult())
}
