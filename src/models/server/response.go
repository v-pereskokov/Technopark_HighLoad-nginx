package server

type Headers map[string]string

type Response struct {
	Status  *Status
	Headers Headers
}

func InitResponse() *Response {
	response := new(Response)

	response.Status = new(Status)
	response.Headers = make(Headers)

	return response
}
