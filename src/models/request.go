package models

import "net/url"

type Request struct {
	method *Method
	url    *url.URL
}

func (request *Request) GetPath() string {
	if request.url != nil {
		return request.url.Path
	}

	return ""
}

func (request *Request) set_path(path string) {
	request.url.Path = path
}
