package constants

var (
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
	INDEX_FILE           string = "/index.html"
	SERVER               string = "VLADDOS Server"
	ERROR_BODY_MIME_TYPE string = "text/html"
	DEFAULT_MIME_TYPE    string = "text/html"
)
