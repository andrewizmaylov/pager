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
	name = flag.String("name", gofakeit.Name(), "Registration Name")
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

	res, err := c.RegisterUser(context.Background(), &pb.RegisterUserRequest{
		Name:     *name,
		Email:    *email,
		Password: *password,
	})

	if err != nil {
		log.Fatalf("Error with user registration %v", err)
	}

	fmt.Printf("User registered: %d", res.Id)
}
