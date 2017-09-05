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

func (server *Server) GetNetwork() string {
	return server.Network
}

func (server *Server) GetProtocol() string {
	return server.Protocol
}

func (server *Server) GetHost() string {
	return server.Host
}

func (server *Server) GetPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = strconv.Itoa(server.Port)
	}

	return port
}
