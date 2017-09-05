package configs

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Server *Server
	Dir    *Dir
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

func (config *Config) GetDir() string {
	config.Dir.FindDir()

	fmt.Println(config.Dir.Path)

	return config.Dir.Path
}
