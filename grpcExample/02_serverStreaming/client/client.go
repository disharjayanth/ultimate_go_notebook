package main

import (
	"context"
	"fmt"
	"io"

	"github.com/disharjayanth/ultimate_go_notebook/grpcExample/02_serverStreaming/multiplepb"
	"google.golang.org/grpc"
)

func main() {
	var num int64
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Error dailing to localhost:50051", err)
		return
	}

	client := multiplepb.NewMultipleServiceClient(conn)

	fmt.Println("Enter num for multiples:")
	fmt.Scanf("%d", &num)

	req := multiplepb.MultipleRequest{
		Num: num,
	}

	resStream, err := client.Multiple(context.Background(), &req)
	if err != nil {
		fmt.Printf("Error from server response: %v\n", err)
		return
	}

	for {
		res, err := resStream.Recv()
		if err == io.EOF {
			fmt.Println("End of response stream.")
			break
		}
		if err != nil {
			fmt.Println("Error while receiving response stream", err)
			return
		}

		fmt.Println(res.GetNum())
	}
}
