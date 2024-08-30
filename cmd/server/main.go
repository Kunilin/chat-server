package main

import (
	"context"
	"fmt"
	desc "gitlab.com/konfka/chat-server/pkg/chat_v1"
	"math/rand"
)

const grpcPort = 50051

type Server struct {
	desc.UnimplementedChatV1Server
}

func (s Server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	fmt.Printf("Creating chat for users: %s", req.GetUsernames())
	return &desc.CreateResponse{Id: rand.Int63()}, nil
}

func (s Server) Delete(ctx context.Context, req *desc.DeleteRequest) error {
	fmt.Printf("Deleting chat with %d id", req.GetId())
	return nil
}

func (s Server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) error {
	fmt.Printf("Sending message from %s", req.GetFrom())
	return nil
}
