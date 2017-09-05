package configs

import (
	"os"
	"strconv"
)

type Config struct {
	server *Server
	dir    *Dir
}

func (config *Config) GetNetwork() string {
	return config.server.Network
}

func (config *Config) GetProtocol() string {
	return config.server.Protocol
}

func (config *Config) GetHost() string {
	return config.server.Host
}

func (config *Config) GetPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = strconv.Itoa(config.server.Port)
	}

	return port
}

func (config *Config) GetDir() string {
	return config.dir.Path
}
