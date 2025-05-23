package book

type ProcessSourceDataResult struct {
	Metadata
	Images
}

type ProcessService interface {
	ProcessSourceData(isbn ISBN, data SourceMultiData) (*CreateParams, error)
}
