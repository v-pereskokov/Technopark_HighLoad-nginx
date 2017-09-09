package server

import (
	"fmt"
	"github.com/vladpereskokov/Technopark_HighLoad-nginx/src/handler"
	modelConfig "github.com/vladpereskokov/Technopark_HighLoad-nginx/src/models/configs"
	"log"
	"net"
	"strconv"
)

type Server struct {
	Network  string
	Protocol string
	Host     string
	Port     string
	IsSetup  bool
}

func CreateServer(config modelConfig.Server) (server Server) {
	server.setNetwork(config.Network)
	server.setProtocol(config.Protocol)
	server.setHost(config.Host)
	server.setPort(strconv.Itoa(config.Port))
	server.setSetup(true)

	return
}

func (server *Server) Start(handle handler.HandlerFunc) {
	if server.IsSetup {
		listener, err := net.Listen(server.Network, ":"+server.Port)
		if err != nil {
			panic("Failed start server: " + err.Error())
		}
		defer listener.Close()

		log.Print("Server started at " + server.Port + " port")

		ch := make(chan net.Conn)

		for i := 0; i < 4; i++ {
			println("Created worker...")
			handle(ch)
		}

		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Error accepting: ", err.Error())

				continue
			}

			ch <- conn
		}
	} else {
		panic("server is not setup")
	}
}

func (server *Server) setNetwork(network string) {
	server.Network = network
}

func (server *Server) setProtocol(protocol string) {
	server.Protocol = protocol
}

func (server *Server) setHost(host string) {
	server.Host = host
}

func (server *Server) setPort(port string) {
	server.Port = port
}

func (server *Server) setSetup(isSetup bool) {
	server.IsSetup = isSetup
}
