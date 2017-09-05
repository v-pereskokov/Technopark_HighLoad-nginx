package constants

import "github.com/vladpereskokov/Technopark_HighLoad-nginx/src/models"

var (
	STATUSES = map[int]models.Status{
		200: {Message: "OK", Code: 200},
		404: {Message: "NOT FOUND", Code: 404},
		400: {Message: "BAD REQUEST", Code: 400},
		403: {Message: "FORBIDDEN", Code: 403},
		405: {Message: "METHOD NOT ALLOWED", Code: 405},
		505: {Message: "HTTP VERSION NOT SUPPORTED", Code: 505},
	}
	IMPLEMENTED_METHODS = []string{
		"GET",
		"HEAD",
	}
	CONTENT_TYPES = map[string]string{
		".css":  "text/css",
		".gif":  "image/gif",
		".html": "text/html",
		".jpeg": "image/jpeg",
		".jpg":  "image/jpeg",
		".js":   "text/javascript",
		".json": "application/json",
		".txt":  "application/text",
		".png":  "image/png",
		".swf":  "application/x-shockwave-flash",
	}
)

const (
	STRING_SEPARATOR     string = "\r\n"
	HEADERS_END          string = STRING_SEPARATOR + STRING_SEPARATOR
	HTTP_VERSION         string = "HTTP/1.1"
	SERVER               string = "IliaBulavintsevServer"
	INDEX_FILE           string = "/index.html"
	ERROR_BODY_MIME_TYPE string = "text/html"
	DEFAULT_MIME_TYPE    string = "text/html"
)
