package server

type Status struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type Statuses []Status
