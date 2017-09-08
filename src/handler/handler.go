package handler

import (
	"fmt"
	modelServer "github.com/vladpereskokov/Technopark_HighLoad-nginx/src/models/server"
	"net"
	"strings"
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
		handler.Connection = <-channel
		handler.handle()
	}
}

func (handler *Handler) handle() {
	handler.readRequest()
	handler.close()
}

func (handler *Handler) readRequest() {
	buffer := make([]byte, 1024)

	_, err := handler.Connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())

		return
	}

	fmt.Printf("%s", buffer)

	handler.Connection.Write(buffer)
	handler.Connection.Write([]byte("\r\n\r\n"))
}

func (handler *Handler) parse_start_string(start_string string) {
	splited_string := strings.Split(start_string, " ")
	if len(splited_string) != 3 {
		handler.Response.Status.Message = "bad_request"
		handler.Response.Status.Code = 400
		return
	}

	fmt.Println("top: ")
	fmt.Printf("%v", splited_string[0])
	fmt.Println("top2:")
}

func (handler *Handler) close() {
	handler.Connection.Close()
}
