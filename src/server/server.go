package server

import (
	"fmt"
	"github.com/vladpereskokov/Technopark_HighLoad-nginx/src/configs"
	"github.com/vladpereskokov/Technopark_HighLoad-nginx/src/handler"
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

func (server *Server) Start(config *configs.Config) {
	listener, err := net.Listen(config.GetNetwork(), ":"+config.GetPort())
	if err != nil {
		panic("Failed start server: " + err.Error())
	}

	defer listener.Close()

	log.Print("Server started at " + config.GetPort() + " port")

	ch := make(chan net.Conn)

	handle := handler.Handler{}

	for i := 0; i < 4; i++ {
		println("Created worker...")
		go handle.Start(ch)
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