package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/andrewizmaylov/pager/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	total = flag.Int64("total", -1, "Amount of users to get back")
	addr  = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Could not connect to client %v", err)
	}

	fmt.Println("Connection OK")

	c := pb.NewPagerClient(conn)

	defer conn.Close()

	log.Printf("Looking for user list \n")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := c.ListRegisteredUsers(ctx, &pb.UserListRequest{
		Total:    *total,
	})

	if err != nil {
		log.Fatalf("User list fetching failed: %v", err)
	}
	for {
		user, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("client.ListFeatures failed: %v", err)
		}
		log.Printf("User: id: %d, name: %s, email: %s, password: %s)", user.GetId(), user.GetName(), user.GetEmail(), user.GetPassword())
	}
}
