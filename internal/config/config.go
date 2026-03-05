package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string        `yaml:"env" env-default:"local"`
	StoragePath string        `yaml:"storage_path" env-required:"true"`
	TokenTTL    time.Duration `yaml:"token_ttl" env-required:"true"`
	GRPC        GRPCConfig    `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port" env-default:"5011"`
	Timeout time.Duration `yaml:"timeout" env-default:"15s"`
}

func Mustload() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("Config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("Config path does not exists: " + path)
	}

	var cnfg Config

	if err := cleanenv.ReadConfig(path, &cnfg); err != nil {
		panic("Failed to read config: " + err.Error())
	}

	return &cnfg
}

// Fetch config path from CL flag or env variable
// Priority flag > env > default
// Default empty string
func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "Path to config")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
