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
	Constants  *modelServer.Constants
	Dir        string
}

type HandlerFunc func(chan net.Conn)

func CreateHandler(config *modelServer.Constants, dir string) (handlerFunc HandlerFunc) {
	handler := Handler{}
	handler.create(config, dir)

	return func(channel chan net.Conn) {
		go handler.start(channel)
	}
}

func (handler *Handler) create(config *modelServer.Constants, dir string) {
	handler.Request = new(modelServer.Request)

	handler.Response = new(modelServer.Response)
	handler.Response.Status = new(modelServer.Status)
	handler.Response.Headers = make(map[string]string)

	handler.Constants = config
	handler.Dir = dir
}

func (handler *Handler) start(channel chan net.Conn) {
	for {
		handler.handle(<-channel)
	}
}

func (handler *Handler) handle(connection net.Conn) {
	buf := make([]byte, 1024)

	_, err := connection.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())

		connection.Close()
		return
	}

	connection.Write([]byte(buf))
	connection.Write([]byte("\r\n\r\n"))
	connection.Close()
}
