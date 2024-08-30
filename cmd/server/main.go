package main

import (
	"context"
	"fmt"
	desc "gitlab.com/konfka/chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"math/rand"
	"net"
)

const grpcPort = 50051

type Server struct {
	desc.UnimplementedChatV1Server
}

func (s Server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	fmt.Printf("Creating chat for users: %s", req.GetUsernames())
	return &desc.CreateResponse{Id: rand.Int63()}, nil
}

func (s Server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	fmt.Printf("Deleting chat with %d id", req.GetId())
	return &emptypb.Empty{}, nil
}

func (s Server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	fmt.Printf("Sending message from %s", req.GetFrom())
	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalln("Failed to listen server: ", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &Server{})

	log.Printf("Server listening at %+v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %e", err)
	}
}
