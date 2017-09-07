package server

type Constants struct {
	Statuses     *Statuses
	Methods      *Methods
	ContentTypes *ContentTypes
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
