package server

import (
	"fmt"
	"github.com/vladpereskokov/Technopark_HighLoad-nginx/src/configs"
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

	for {
		// Listen for an incoming connection.
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	// Send a response back to person contacting us.
	conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	conn.Close()
}
