package book

type ProcessSourceDataResult struct {
	Metadata
	Images
}

type ProcessService interface {
	ProcessSourceData(data SourceMultiData) (*CreateParams, error)
}
