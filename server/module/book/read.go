package book

type ReadService interface {
	FromISBN(isbn ISBN, process bool) (*Entity, error)
}
