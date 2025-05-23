package book

type ReadService interface {
	FromISBN(isbn string, process bool) (*Entity, error)
}
