package server

import (
	"fmt"
	"github.com/vladpereskokov/Technopark_HighLoad-nginx/src/handler"
	"github.com/vladpereskokov/Technopark_HighLoad-nginx/src/models"
	"log"
	"net"
	"os"
)

type Server struct {
	network  string
	protocol string
	host     string
	port     string
}

func (server *Server) Start(config *models.Config) {
	serverConf := config.GetServer()

	listener, err := net.Listen(serverConf.GetNetwork(), ":"+serverConf.GetPort())
	if err != nil {
		panic("Failed start server: " + err.Error())
	}

	defer listener.Close()

	log.Print("Server started at " + serverConf.GetPort() + " port")

	ch := make(chan net.Conn)

	handle := handler.Handler{}

	for i := 0; i < 4; i++ {
		go handle.Start(ch)
		println("Created worker...")
	}

	for {
		// Listen for an incoming connection.
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		ch <- conn
	}
}
