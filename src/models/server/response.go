package server

import (
	"bufio"
	"fmt"
	"os"
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

func (response *Response) IsOk() bool {
	return response.Status.Code == 200
}

func (response Response) GetOkBody(path string) (*bufio.Reader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return bufio.NewReader(file), nil
}

func (response Response) GetErrorBody() {

}

func InitResponse() *Response {
	response := new(Response)

	response.Status = new(Status)
	response.Headers = make(Headers)

	return response
}
