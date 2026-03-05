package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/andrewizmaylov/pager/internal/app"
	"github.com/andrewizmaylov/pager/internal/config"
)

func main() {
	cnfg := config.Mustload()

	application := app.New(cnfg.GRPC.Port, cnfg.StoragePath, cnfg.TokenTTL)

	go application.GRPCSrv.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	s := <-stop
	log.Printf("Signal recieved %s", s)

	time.Sleep(4 * time.Second)
	application.GRPCSrv.Stop()
	log.Printf("Apllication shutdown")
}
