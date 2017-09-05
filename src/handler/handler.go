package handler

import (
	"fmt"
	"github.com/vladpereskokov/Technopark_HighLoad-nginx/src/models"
	"net"
)

type Handler struct {
	Connection net.Conn
	Request    *models.Request
	Response   *models.Response
	Dir        string
}

func (handler *Handler) Start(channel chan net.Conn) {
	for {
		conn := <-channel
		buf := make([]byte, 1024)

		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}

		conn.Write([]byte("Message received."))
		conn.Close()
	}
}
