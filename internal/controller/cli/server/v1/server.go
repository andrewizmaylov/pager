package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/andrewizmaylov/pager/proto/v1"
	"github.com/andrewizmaylov/pager/internal/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var cnfg = config.Mustload()

var (
	port = flag.Int("port", cnfg.GRPC.Port, "The server port")
)

var registeredUsers = make(map[string]*pb.UserResponse, 100)

type server struct {
	pb.UnimplementedPagerServer
}

var lastUserId int32 = 1

func (s *server) RegisterUser(ctx context.Context, in *pb.RegisterUserRequest) (*pb.UserResponse, error) {
	out := &pb.UserResponse{
		Id:    lastUserId,
		Name:  in.GetName(),
		Email: in.GetEmail(),
		Password: in.GetPassword(),
	}

	registeredUsers[out.Email] = out

	fmt.Printf("ID: %d, Name: %s, Email: %s\n", lastUserId, in.GetName(), in.GetEmail())

	lastUserId++

	return out, nil
}

func (s *server) LoginUser(ctx context.Context, in *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	user, ok := registeredUsers[in.Email]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "user not found: %s", in.Email)
	}
	if user.Password != in.Password {
		return nil, status.Errorf(codes.InvalidArgument, "check provided password: %s", in.Password)
	}

	out := &pb.LoginUserResponse{
		Id:    user.GetId(),
		Token:  in.GetEmail(),
	}

	log.Printf("User: id: %d, name: %s, email: %s, password: %s)", user.GetId(), user.GetName(), user.GetEmail(), user.GetPassword())
	return out, nil
}

func (s *server) ListRegisteredUsers(in *pb.UserListRequest, stream pb.Pager_ListRegisteredUsersServer) error {
	for _, user := range registeredUsers {
		if err := stream.Send(user); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	port := fmt.Sprintf(":%d", *port)
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to listen port %s", port)
	}

	s := grpc.NewServer()
	pb.RegisterPagerServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
