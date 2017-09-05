package models

type Headers map[string]string

type Response struct {
	Status  *Status
	Headers Headers
}
