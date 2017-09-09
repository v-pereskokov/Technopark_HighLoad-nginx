package server

import (
	"fmt"
	"io/ioutil"
)

type Headers map[string]string

type Response struct {
	Status  *Status
	Headers Headers
}

func (response *Response) SetStatus(code int, statuses *Statuses) {
	for _, value := range *statuses {
		if value.Code == code {
			response.Status.SetStatus(value.Message, value.Code)
		}
	}
}

func (response *Response) SetHeader(cType string, value string) {
	response.Headers[cType] = value
}

func (response *Response) IsOk() bool {
	return response.Status.Code == 200
}

func (response Response) GetOkBody(path string) ([]byte, error) {
	fmt.Print("path: ")
	fmt.Println(path)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (response Response) GetErrorBody(message string) string {
	return "<html><body><h1>" + message + "</h1></body></html>"
}

func InitResponse() *Response {
	response := new(Response)

	response.Status = new(Status)
	response.Headers = make(Headers)

	return response
}
