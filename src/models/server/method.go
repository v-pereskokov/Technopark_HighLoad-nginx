package server

type Method struct {
	Type string `json:"type"`
}

func (method *Method) SetMethod(value string) {
	method.Type = value
}

func (method *Method) GetMethod() string {
	return method.Type
}

type Methods []Method
