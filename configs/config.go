package configs

import (
	"os"
	"strconv"
)

type Config struct {
	Network  string `json:"network"`
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

func (config *Config) GetNetwork() string {
	return config.Network
}

func (config *Config) GetProtocol() string {
	return config.Protocol
}

func (config *Config) GetHost() string {
	return config.Host
}

func (config *Config) GetPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = strconv.Itoa(config.Port)
	}

	return port
}
