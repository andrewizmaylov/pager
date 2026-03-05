package grpcapp

import (
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/andrewizmaylov/pager/internal/grpc/server"
	"google.golang.org/grpc"
)

type App struct {
	gRPCServer *grpc.Server
	port       string
}

func New(port int) *App {
	gRPCServer := grpc.NewServer()

	server.Register(gRPCServer)

	return &App{
		gRPCServer: gRPCServer,
		port:       ":" + strconv.Itoa(port),
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grcpapp.Run"

	l, err := net.Listen("tcp", a.port)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Printf("%s: Starting grpcapp server\n", op)

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "grcpapp.Stop"

	log.Printf("%s: Stop grpcapp server\n", op)

	a.gRPCServer.GracefulStop()
}
