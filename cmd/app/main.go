package main

import (
	"fmt"
	"github.com/andrewizmaylov/pager/internal/config"
)

func main() {
	cnfg := config.Mustload()

	fmt.Printf("Config: %d\n", cnfg.GRPC.Port)
}
