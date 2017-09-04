package server

import (
	"github.com/vladpereskokov/Technopark_HighLoad-nginx/configs"
	"log"
	"net"
)

type Server struct {
	protocol string
	host     string
	port     string
}

func (server *Server) Start(config *configs.Config) {
	_, err := net.Listen("tcp", ":"+config.GetPort())
	if err != nil {
		panic("Failed start server: " + err.Error())
	}
	log.Print("Server started at " + config.GetPort() + " port")
}
