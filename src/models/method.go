package models

type Method struct {
	value string
}

func (method *Method) SetMethod(value string) {
	method.value = value
}

func (method *Method) GetMethod() string {
	return method.value
}
