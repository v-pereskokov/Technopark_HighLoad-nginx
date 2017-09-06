package server

import (
	"fmt"
	"github.com/vladpereskokov/Technopark_HighLoad-nginx/src/handler"
	modelConfig "github.com/vladpereskokov/Technopark_HighLoad-nginx/src/models/configs"
	"log"
	"net"
	"os"
	"strconv"
)

type Server struct {
	Network  string
	Protocol string
	Host     string
	Port     string
	IsSetup  bool
}

func (server *Server) CreateServer(config modelConfig.Server) {
	server.setNetwork(config.Network)
	server.setProtocol(config.Protocol)
	server.setHost(config.Host)
	server.setPort(strconv.Itoa(config.Port))
	server.setSetup(true)
}

func (server *Server) Start() {
	if server.IsSetup {
		listener, err := net.Listen(server.Network, ":"+server.Port)
		if err != nil {
			panic("Failed start server: " + err.Error())
		}

		defer listener.Close()

		log.Print("Server started at " + server.Port + " port")

		ch := make(chan net.Conn)

		handle := handler.Handler{}

		for i := 0; i < 4; i++ {
			go handle.Start(ch)
			println("Created worker...")
		}

		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Error accepting: ", err.Error())
				os.Exit(1)
			}

			ch <- conn
		}
	} else {
		panic("hop hey lalaley")
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
