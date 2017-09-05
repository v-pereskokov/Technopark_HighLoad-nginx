package models

import "net/url"

type Request struct {
	method *Method
	url    *url.URL
}
