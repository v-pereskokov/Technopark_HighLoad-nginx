package server

type Constants struct {
	Statuses     []Status
	Methods      []Method
	ContentTypes []ContentType
}

func (constants *Constants) GetContentTypes() []ContentType {
	return constants.ContentTypes
}

func (constants *Constants) GetMethods() []Method {
	return constants.Methods
}

func (constants *Constants) GetStatuses() []Status {
	return constants.Statuses
}
