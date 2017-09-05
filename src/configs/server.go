package configs

import (
	"os"
	"strconv"
)

type Server struct {
	Network  string `json:"network"`
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

func (config *Config) GetNetwork() string {
	return config.Server.Network
}

func (config *Config) GetProtocol() string {
	return config.Server.Protocol
}

func (config *Config) GetHost() string {
	return config.Server.Host
}

func (config *Config) GetPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = strconv.Itoa(config.Server.Port)
	}

	return port
}
