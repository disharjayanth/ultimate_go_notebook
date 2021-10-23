package main

import (
	"fmt"
	"io"
	"net"

	"github.com/disharjayanth/ultimate_go_notebook/grpcExample/04_biStreaming/maxpb"
	"google.golang.org/grpc"
)

type server struct {
	maxpb.UnimplementedMaxServiceServer
}

func (s *server) Max(stream maxpb.MaxService_MaxServer) error {
	max := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// end of client stream
			return nil
		}
		if err != nil {
			fmt.Println("error receving client stream:", err)
			return err
		}

		num := req.GetNum()
		if num > int64(max) {
			max = int(num)
			res := maxpb.MaxResponse{
				Res: num,
			}
			if err = stream.Send(&res); err != nil {
				fmt.Println("Error sending response to client:", err)
				return err
			}
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		fmt.Println("Error tcp listening at port 50051", err)
		return
	}

	grpcServer := grpc.NewServer()

	maxpb.RegisterMaxServiceServer(grpcServer, &server{})

	fmt.Println("Listening server @port:50051")
	if err = grpcServer.Serve(lis); err != nil {
		fmt.Println("Error serving server at port:50051", err)
		return
	}
}
