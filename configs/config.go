package configs

import (
	"os"
	"strconv"
)

type Config struct {
	Port int `json:"port"`
}

func (config *Config) GetPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = strconv.Itoa(config.Port)
	}

	return port
}
