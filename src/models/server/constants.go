package server

type Constants struct {
	ContentTypes *ContentTypes
	Methods      *Methods
	Statuses     *Statuses
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
