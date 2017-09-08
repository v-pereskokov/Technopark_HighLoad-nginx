package handler

import (
	"bufio"
	"bytes"
	"fmt"
	modelServer "github.com/vladpereskokov/Technopark_HighLoad-nginx/src/models/server"
	"net"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type Handler struct {
	Connection net.Conn
	Request    *modelServer.Request
	Response   *modelServer.Response
	Constants  *modelServer.Constants
	Dir        string
}

type HandlerFunc func(chan net.Conn)

// ?
func CreateHandler(config *modelServer.Constants, dir string) (handlerFunc HandlerFunc) {
	handler := Handler{}
	handler.create(config, dir)

	return func(channel chan net.Conn) {
		go handler.start(channel)
	}
}

func (handler *Handler) create(config *modelServer.Constants, dir string) {
	handler.Request = new(modelServer.Request)
	handler.Request.Method = new(modelServer.Method)
	handler.Request.Url = new(url.URL)

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
	handler.process_request()
	handler.write_response()
	handler.closeConn()
}

func (handler *Handler) readRequest() {
	buffer := make([]byte, 1024)

	_, err := handler.Connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())

		return
	}

	raw_request := string(buffer[:bytes.Index(buffer, []byte{0})])
	start_string := strings.Split(raw_request, "\r\n")[0]
	fmt.Println(start_string)
	handler.parse_start_string(start_string)
}

func (handler *Handler) parse_start_string(start_string string) {
	splited_string := strings.Split(start_string, " ")
	if len(splited_string) != 3 {
		handler.Response.Status.Message = "bad_request"
		handler.Response.Status.Code = 400

		return
	}
	handler.Request.Method.Type = splited_string[0]
	parsed_url, err := url.Parse(splited_string[1])
	if err != nil || !strings.HasPrefix(splited_string[2], "HTTP/") {
		handler.Response.Status.Message = "bad_request"
		handler.Response.Status.Code = 400
	}
	handler.Request.Url = parsed_url
}

func (handler *Handler) process_request() {
	if handler.Response.Status.Code != 200 {
		handler.set_content_headers(nil)
		return
	}

	if !handler.Constants.Methods.Contains(handler.Request.Method.Type) {
		handler.Response.Status.Message = "method_not_allowed"
		handler.Response.Status.Code = 405
	} else {
		handler.preprocess_path()
	}
}

//preproccess path and check file errors
func (handler *Handler) preprocess_path() {
	handler.Request.SetPath(handler.Dir + handler.Request.GetPath())
	file_info := handler.check_path(false)

	if file_info != nil && file_info.IsDir() {
		handler.Request.SetPath(handler.Request.GetPath() + "/index.html")
		file_info = handler.check_path(true)
	}

	handler.set_content_headers(file_info)
}

func (handler *Handler) check_path(is_dir bool) os.FileInfo {
	request_path := handler.Request.GetPath()
	clear_path := path.Clean(request_path)
	handler.Request.SetPath(clear_path)
	info, err := os.Stat(request_path)
	if err != nil {

		if os.IsNotExist(err) && !is_dir {
			handler.Response.Status.Message = "not_found"
			handler.Response.Status.Code = 404
		} else {
			handler.Response.Status.Message = "forbidden"
			handler.Response.Status.Code = 403
		}
	} else if !strings.Contains(clear_path, handler.Dir) {
		handler.Response.Status.Message = "forbidden"
		handler.Response.Status.Code = 403
	}
	return info
}

func (handler *Handler) set_content_headers(info os.FileInfo) {
	if handler.Response.Status.Code == 200 {
		handler.Response.Headers["Content-Length"] = strconv.Itoa(int(info.Size()))
		handler.Response.Headers["Content-Type"] = handler.get_content_type()
	} else {
		handler.Response.Headers["Content-Length"] = strconv.Itoa(len(handler.get_error_body()))
		handler.Response.Headers["Content-Type"] = "text/html"
	}
}

func (handler *Handler) get_error_body() string {
	body := "<html><body><h1>"
	body += handler.Response.Status.Message
	body += "</h1></body></html>"
	return body
}

func (handler *Handler) get_content_type() string {
	extension := ""

	request_path := handler.Request.GetPath()
	last_dot := strings.LastIndex(request_path, ".")

	if last_dot >= 0 {
		extension = request_path[last_dot:]
	}

	val := handler.Constants.ContentTypes.GetType(extension)
	if len(val) != 0 {
		return val
	} else {
		return "text/html"
	}
}

func (handler Handler) write_response() {
	handler.write_string("HTTP/1.1" + " " + handler.Response.Status.Message)
	handler.write_headers()
	handler.write_string("") // empty string after headers
	if handler.Request.Method.Type != "HEAD" {
		handler.write_body()
	}
	fmt.Println(handler.Request.Method.Type, " ", handler.Request.GetPath(), " ", handler.Response.Status.Code)
}

func (handler *Handler) write_string(str string) {
	handler.Connection.Write([]byte(str + "\r\n"))
}

func (handler *Handler) write_body() {
	if handler.Response.Status.Code == 200 {
		handler.write_ok_body()
	} else {
		handler.write_error_body()
	}
}

func (handler *Handler) write_ok_body() {
	file, err := os.Open(handler.Request.GetPath())
	if err != nil {
		fmt.Println("Can't open file ", handler.Request.GetPath())
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	_, read_err := reader.WriteTo(handler.Connection)
	if read_err != nil {
		fmt.Println("Some error on read or write file ", handler.Request.GetPath())
	}
}

func (handler *Handler) write_error_body() {
	body := []byte(handler.get_error_body())
	handler.Connection.Write(body)
}

func (handler Handler) write_headers() {
	handler.write_common_headers()
	handler.write_specific_headers()
}

func (handler Handler) write_common_headers() {
	handler.write_string("Date: " + time.Now().String())
	handler.write_string("Server: " + "Vladdos")
	handler.write_string("Connection: close")
}

func (handler Handler) write_specific_headers() {
	for key, value := range handler.Response.Headers {
		handler.write_string(key + ": " + value)
	}
}

func (handler *Handler) closeConn() {
	handler.Connection.Close()
}
