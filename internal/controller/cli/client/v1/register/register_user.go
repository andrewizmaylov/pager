package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode/utf8"

	pb "github.com/andrewizmaylov/pager/proto/v1"
	"github.com/brianvoe/gofakeit/v6"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var fakeName string

func getName() string {
	fakeName = gofakeit.Name()
	return fakeName
}

func getEmail() string {
	return strings.ReplaceAll(strings.ToLower(fakeName), " ", ".") + "@yost.biz"
}

var (
	port     = flag.Int("port", 8080, "TCP port for connection")
	name     = flag.String("name", "", "Registration Name")
	email    = flag.String("email", "", "Registration Email")
	password = flag.String("password", "123456", "Password")
)

func main() {
	flag.Parse()

	checkNameAndEmail()

	conn, err := grpc.NewClient("localhost:"+strconv.Itoa(*port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Could not connect to client %v\n", err)
	}

	c := pb.NewPagerClient(conn)

	defer conn.Close()

	res, err := c.RegisterUser(context.Background(), &pb.RegisterUserRequest{
		Name:     *name,
		Email:    *email,
		Password: *password,
	})

	if err != nil {
		log.Fatalf("Error with user registration %v\n", err)
	}

	fmt.Printf("User registered: %d\n", res.Id)
}

func checkNameAndEmail() {
	if utf8.RuneCountInString(*name) > 0 {
		fakeName = *name
	}

	if *name == "" {
		*name = getName()
	}
	if *email == "" {
		*email = getEmail()
	}
}
