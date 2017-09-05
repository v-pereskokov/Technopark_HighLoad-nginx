package handler

import (
	"fmt"
	"net"
)

type Handler struct {
}

func (handler *Handler) Start(channel chan net.Conn) {
	for {
		conn := <-channel
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
}
