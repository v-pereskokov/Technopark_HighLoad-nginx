package handler

import (
	"fmt"
	"github.com/vladpereskokov/Technopark_HighLoad-nginx/src/constants"
	modelServer "github.com/vladpereskokov/Technopark_HighLoad-nginx/src/models/server"
	"github.com/vladpereskokov/Technopark_HighLoad-nginx/src/utils"
	"log"
	"net"
	"strings"
	"time"
)

const HTTP_CONSTANTS_CONFIG = "configs/http.json"

type Handler struct {
	Connection net.Conn
	Request    *modelServer.Request
	Response   *modelServer.Response
	Constants  *modelServer.Constants
	Dir        string
}

type HandlerFunc func(chan net.Conn)

func CreateHandler(dir string) (handlerFunc HandlerFunc) {
	httpConstantsConfig := new(modelServer.Constants)

	err := utils.FromFile(HTTP_CONSTANTS_CONFIG, &httpConstantsConfig)
	if err != nil {
		log.Panicf("can not init http config: %v", err)
	}

	return func(channel chan net.Conn) {
		handler := Handler{}
		handler.create(httpConstantsConfig, dir)

		go handler.start(channel)
	}
}

func (handler *Handler) create(config *modelServer.Constants, dir string) {
	handler.Request = modelServer.InitRequest()
	handler.Response = modelServer.InitResponse()

	handler.Response.SetStatus(200, config.Statuses)

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
	handler.writeResponse()
	handler.closeConn()
}

func (handler *Handler) readRequest() {
	buffer := make([]byte, 1024)

	_, err := handler.Connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())

		return
	}

	handler.parseRequest(string(buffer))
}

func (handler *Handler) parseRequest(query string) {
	splitedQuery := strings.Split(query, "\r\n")[0]
	queryParts := strings.Split(splitedQuery, " ")

	for _, value := range queryParts {
		handler.Connection.Write([]byte(value))
		handler.Connection.Write([]byte("\r\n"))
	}

	if len(queryParts) != 3 {
		handler.Response.SetStatus(400, handler.Constants.Statuses)

		return
	}
}

func (handler Handler) writeResponse() {
	handler.write(constants.HTTP_VERSION + " " + handler.Response.Status.Message)
	handler.writeHeader()
}

func (handler Handler) writeHeader() {
	handler.writeCommonHeaders()
	handler.writeSpecificHeaders()
	handler.write("")
	if handler.Request.Method.Type != "HEAD" {
		handler.writeBody()
	}
}

func (handler Handler) writeCommonHeaders() {
	handler.write("Date: " + time.Now().String())
	handler.write("Server: " + constants.SERVER)
	handler.write("Connection: close")
}

func (handler Handler) writeSpecificHeaders() {
	for key, value := range handler.Response.Headers {
		handler.write(key + ": " + value)
	}
}

func (handler Handler) writeBody() {
	if handler.Response.IsOk() {
		handler.writeOkBody()
	} else {
		handler.Response.GetErrorBody()
	}
}

func (handler Handler) writeOkBody() {
	reader, err := handler.Response.GetOkBody(handler.Request.GetPath())
	if err != nil {
		fmt.Println("Can't open file ", handler.Request.GetPath())
	}

	_, err = reader.WriteTo(handler.Connection)
	if err != nil {
		fmt.Println("Some error on read or write file ", handler.Request.GetPath())
	}
}

func (handler Handler) write(content string) {
	handler.Connection.Write([]byte(content + "\r\n"))
}

func (handler *Handler) closeConn() {
	handler.Connection.Close()
}
