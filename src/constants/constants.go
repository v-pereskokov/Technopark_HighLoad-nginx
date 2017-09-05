package constants

import "github.com/vladpereskokov/Technopark_HighLoad-nginx/src/models"

var (
	STATUSES = map[string]models.Status{
		"ok":                 {Message: "200 OK", Code: 200},
		"not_found":          {Message: "404 NOT FOUND", Code: 404},
		"bad_request":        {Message: "400 BAD REQUEST", Code: 400},
		"forbidden":          {Message: "403 FORBIDDEN", Code: 403},
		"method_not_allowed": {Message: "405 METHOD NOT ALLOWED", Code: 405},
		"not_supports":       {Message: "505 HTTP VERSION NOT SUPPORTED", Code: 505},
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
