package server

type Status struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (status *Status) SetStatus(message string, code int) {
	status.Message = message
	status.Code = code
}

type Statuses []Status
