package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"
	"unicode/utf8"

	pb "github.com/andrewizmaylov/pager/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


var (
	port     = flag.Int("port", 8080, "TCP port for connection")
	email    = flag.String("email", "", "Registration Email")
	password = flag.String("password", "123456", "Password")
)

func main() {
	flag.Parse()

	if utf8.RuneCountInString(*email) == 0 {
		log.Fatalf("Empty email")
	}

	conn, err := grpc.NewClient("localhost:"+strconv.Itoa(*port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Could not connect to client %v", err)
	}

	c := pb.NewPagerClient(conn)

	defer conn.Close()

	res, err := c.LoginUser(context.Background(), &pb.LoginUserRequest{
		Email:    *email,
		Password: *password,
	})

	if err != nil {
		log.Fatalf("User login error: %v", err)
	}

	fmt.Printf("User succesfuly logged Id: %d, token %s\n", res.GetId(), res.GetToken())
}
