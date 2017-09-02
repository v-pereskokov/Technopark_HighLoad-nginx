package server

import (
	"fmt"
	"net"
)

type Server struct {
	port string
	host string
}

func (s *Server) Create(host string, port string) {
	s.port = port
	s.host = host
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", s.host+":"+s.port)
	if err != nil {
		fmt.Println("Start server failed ", err)
		return
	} else {
		fmt.Println("Server start ", listener.Addr())
	}
	for {
		_, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("CREATE Handler!!\n\n")
	}
}
