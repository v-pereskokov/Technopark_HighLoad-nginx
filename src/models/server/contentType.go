package server

type ContentType struct {
	Expansion string `json:"expansion"`
	Type      string `json:"type"`
}

type ContentTypes []ContentType

func (types *ContentTypes) GetType(expansion string) (cType string) {
	for _, value := range *types {
		if value.Expansion == expansion {
			return value.Type
		}
	}

	return ""
}
