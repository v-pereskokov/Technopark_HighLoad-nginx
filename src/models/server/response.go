package server

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

func InitResponse() *Response {
	response := new(Response)

	response.Status = new(Status)
	response.Headers = make(Headers)

	return response
}
