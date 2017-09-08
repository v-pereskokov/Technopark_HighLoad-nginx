package server

import "net/url"

type Request struct {
	Method *Method
	Url    *url.URL
}

func (request *Request) GetPath() string {
	if request.Url != nil {
		return request.Url.Path
	}

	return ""
}

func (request *Request) SetPath(path string) {
	request.Url.Path = path
}

func InitRequest() *Request {
	request := new(Request)
	request.Method = new(Method)
	request.Url = new(url.URL)

	return request
}
