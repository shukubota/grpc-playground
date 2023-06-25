package handler

import (
	"context"
	"fmt"
	example "github.com/shukubota/grpc-playground/gen/go/proto"
)

type ExampleServer struct {
	example.UnimplementedExampleServer
}

func NewExampleAPIServer() *ExampleServer {
	return &ExampleServer{}
}

func (s *ExampleServer) GetMessage(ctx context.Context, r *example.GetMessageRequest) (*example.GetMessageResponse, error) {
	fmt.Println("GetMessage request")
	return &example.GetMessageResponse{
		Message: "OK",
	}, nil
}
