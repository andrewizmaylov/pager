package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/andrewizmaylov/pager/proto/v1"
	"github.com/brianvoe/gofakeit/v6"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	email = flag.String("email", gofakeit.Email(), "Registration Email")
	password = flag.String("password", "123456", "Password")
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

	res, err := c.LoginUser(context.Background(), &pb.LoginUserRequest{
		Email:    *email,
		Password: *password,
	})

	if err != nil {
		log.Fatalf("User login error: %v", err)
	}

	fmt.Printf("User succesfuly logged In: %d", res.Id)
}
