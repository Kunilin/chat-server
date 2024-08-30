package main

import (
	"context"
	desc "gitlab.com/konfka/chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const address = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %e", err)
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Printf("Failed to close connection: %e", err)
		}
	}()

	client := desc.NewChatV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	usernames := []string{"Pidor", "Huesos"}

	r, err := client.Create(ctx, &desc.CreateRequest{Usernames: usernames})
	if err != nil {
		log.Fatalf("Failed to create chat: %e", err)
	}

	log.Printf("Chat info: %+v", r)
}
