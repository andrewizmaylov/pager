package server

import (
	pagerv01 "github.com/andrewizmaylov/pager/proto/v1"
	"google.golang.org/grpc"
)


type serverAPI struct {
	pagerv01.UnimplementedPagerServer
}

func Register(gRPC *grpc.Server) {
	pagerv01.RegisterPagerServer(gRPC, &serverAPI{})
}
