package server

type Constants struct {
	Statuses     *Statuses     `json:"status"`
	Methods      *Methods      `json:"http_methods"`
	ContentTypes *ContentTypes `json:"content_type"`
}

func (constants *Constants) GetContentTypes() *ContentTypes {
	return constants.ContentTypes
}

func (constants *Constants) GetMethods() *Methods {
	return constants.Methods
}

func (constants *Constants) GetStatuses() *Statuses {
	return constants.Statuses
}
