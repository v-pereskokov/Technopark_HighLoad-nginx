package server

type ContentType struct {
	Expansion string `json:"expansion"`
	Type      string `json:"type"`
}

type ContentTypes []ContentType
