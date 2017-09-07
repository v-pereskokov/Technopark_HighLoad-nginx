package handler

import (
	"fmt"
	modelServer "github.com/vladpereskokov/Technopark_HighLoad-nginx/src/models/server"
	"net"
)

type Handler struct {
	Connection net.Conn
	Request    *modelServer.Request
	Response   *modelServer.Response
	Dir        string
}

func CreateHandler(dir string) (handlerFunc HandlerFunc) {
	handler := Handler{}
	handler.create(dir)

	return func(channel chan net.Conn) {
		go handler.handle(channel)
	}
}

func (handler *Handler) create(dir string) {
	handler.Request = new(modelServer.Request)

	handler.Response = new(modelServer.Response)
	handler.Response.Status = new(modelServer.Status)
	handler.Response.Headers = make(map[string]string)

	handler.Dir = dir
}

func (handler *Handler) handle(channel chan net.Conn) {
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
