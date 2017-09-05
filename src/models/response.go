package models

import "github.com/vladpereskokov/Technopark_HighLoad-nginx/src/constants"

type Headers map[string]string

type Response struct {
	Status  *Status
	Headers Headers
}

func (r *Response) set_status(status string) {
	if _, ok := constants.STATUSES[status]; ok {
		*r.Status = constants.STATUSES[status]
	}
}

func (r *Response) is_ok() bool {
	return *r.Status == constants.STATUSES["ok"]
}
