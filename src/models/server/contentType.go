package server

type ContentType struct {
	Extension string `json:"extension"`
	Type      string `json:"type"`
}

type ContentTypes []ContentType

func (types *ContentTypes) GetType(extension string) (cType string) {
	for _, value := range *types {
		if value.Extension == extension {
			return value.Type
		}
	}

	return ""
}
