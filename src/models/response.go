package models

import "github.com/vladpereskokov/Technopark_HighLoad-nginx/src/constants"

type Headers map[string]string

type Response struct {
	Status  *Status
	Headers Headers
}

func (r *Response) SetStatus(status int) {
	if _, ok := constants.STATUSES[status]; ok {
		*r.Status = constants.STATUSES[status]
	}
}

func (r *Response) IsOk() bool {
	return *r.Status == constants.STATUSES[200]
}
