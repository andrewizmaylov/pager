package app

import (
	"time"

	grpcapp "github.com/andrewizmaylov/pager/internal/app/grpc"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {

	// TODO init storage etc

	grcpApp := grpcapp.New(grpcPort)

	return &App {
		GRPCSrv: grcpApp,
	}
}
