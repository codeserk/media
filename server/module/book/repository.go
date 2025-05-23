package book

type FilterParams struct {
	Title      string
	TitleExact string
	ISBN       string
}

type CreateParams struct {
	Metadata
	Sources SourceMultiData
	Images
}

type UpdateParams struct {
	Metadata
	Images
}

type Repository interface {
	GetById(id string) (*Entity, error)
	GetByISBN(isbn ISBN) (*Entity, error)

	Create(params CreateParams) (*Entity, error)

	Update(id string, params UpdateParams) (*Entity, error)

	Archive(id string) (*Entity, error)
	Delete(id string) error
}
