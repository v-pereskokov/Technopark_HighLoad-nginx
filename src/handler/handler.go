package handler

import (
	"fmt"
	"github.com/vladpereskokov/Technopark_HighLoad-nginx/src/constants"
	modelServer "github.com/vladpereskokov/Technopark_HighLoad-nginx/src/models/server"
	"github.com/vladpereskokov/Technopark_HighLoad-nginx/src/utils"
	"log"
	"net"
	"net/url"
	"os"
	"path"
	"strconv"
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
	handler.requestHandle()
}

func (handler *Handler) parseRequest(query string) {
	splitedQuery := strings.Split(query, "\r\n")[0]
	queryParts := strings.Split(splitedQuery, " ")

	if len(queryParts) != 3 {
		handler.Response.SetStatus(400, handler.Constants.Statuses)

		return
	}

	handler.Request.Method.SetMethod(queryParts[0])
	parsed_url, err := url.Parse(queryParts[1])

	if err != nil || !strings.HasPrefix(queryParts[2], "HTTP/") {
		handler.Response.SetStatus(400, handler.Constants.Statuses)
	}

	handler.Request.Url = parsed_url
}

func (handler *Handler) requestHandle() {
	if !handler.Response.IsOk() {
		handler.setContentHeaders(nil)
		return
	}

	if !handler.Constants.Methods.Contains(handler.Request.Method.Type) {
		handler.Response.SetStatus(405, handler.Constants.Statuses)
	} else {
		handler.preProcessPath()
	}
}

func (handler *Handler) preProcessPath() {
	handler.Request.SetPath(handler.Dir + handler.Request.GetPath())
	file := handler.checkPath(false)

	if file != nil && file.IsDir() {
		handler.Request.SetPath(handler.Request.GetPath() + constants.INDEX_FILE)
		file = handler.checkPath(true)
	}

	handler.setContentHeaders(file)
}

func (handler *Handler) checkPath(is_dir bool) os.FileInfo {
	requestPath := handler.Request.GetPath()

	clearPath := path.Clean(requestPath)
	handler.Request.SetPath(clearPath)

	info, err := os.Stat(requestPath)
	if err != nil {
		if os.IsNotExist(err) && !is_dir {
			handler.Response.SetStatus(404, handler.Constants.Statuses)
		} else {
			handler.Response.SetStatus(403, handler.Constants.Statuses)
		}
	} else if !strings.Contains(clearPath, handler.Dir) {
		handler.Response.SetStatus(403, handler.Constants.Statuses)
	}

	return info
}

func (handler *Handler) setContentHeaders(info os.FileInfo) {
	if handler.Response.IsOk() {
		handler.Response.SetHeader("Content-Length", strconv.Itoa(int(info.Size())))
		handler.Response.SetHeader("Content-Type", handler.getContentType())
	} else {
		handler.Response.SetHeader("Content-Length",
			strconv.Itoa(len(handler.Response.GetErrorBody(handler.Response.Status.Message))))
		handler.Response.SetHeader("Content-Type", constants.ERROR_BODY_MIME_TYPE)
	}
}

func (handler *Handler) getContentType() string {
	extension := ""

	requestPath := handler.Request.GetPath()
	lastDot := strings.LastIndex(requestPath, ".")

	if lastDot >= 0 {
		extension = requestPath[lastDot:]
	}

	cType := handler.Constants.ContentTypes.GetType(extension)

	if len(cType) != 0 {
		return cType
	} else {
		return constants.DEFAULT_MIME_TYPE
	}
}

func (handler Handler) writeResponse() {
	handler.write(constants.HTTP_VERSION + " " +
		strconv.Itoa(handler.Response.Status.Code) + " " + handler.Response.Status.Message)
	handler.writeHeader()

	handler.write("")
	if handler.Request.Method.Type != "HEAD" {
		handler.writeBody()
	}

	handler.write("\r\n")
}

func (handler Handler) writeHeader() {
	handler.writeCommonHeaders()
	handler.writeSpecificHeaders()
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
		handler.writeErrorBody()
	}
}

func (handler Handler) writeOkBody() {
	file, err := handler.Response.GetOkBody(handler.Request.GetPath())
	if err != nil {
		fmt.Println("Can't open file ", handler.Request.GetPath())
		return
	}

	handler.write(string(file))
}

func (handler Handler) writeErrorBody() {
	body := handler.Response.GetErrorBody(handler.Response.Status.Message)

	handler.write(body)
}

func (handler Handler) write(content string) {
	handler.Connection.Write([]byte(content + "\r\n"))
}

func (handler *Handler) closeConn() {
	handler.Connection.Close()
}
