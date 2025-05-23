package read

import (
	"fmt"
	"media/module/book"
)

func (s *service) FromISBN(isbn string, process bool) (*book.Entity, error) {
	res, err := s.fromISBNInDB(isbn)
	if err != nil {
		return nil, fmt.Errorf("in db: %v", err)
	}
	if res != nil {
		return res, nil
	}

	res, err = s.fromISBNInSources(isbn)
	if err != nil {
		return nil, fmt.Errorf("in sources: %v", err)
	}

	return res, nil
}

func (s *service) fromISBNInDB(isbn string) (*book.Entity, error) {
	book, err := s.repository.GetByISBN(isbn)
	if err != nil {
		return nil, fmt.Errorf("from repository: %v", err)
	}

	return book, nil
}

func (s *service) fromISBNInSources(isbn string) (*book.Entity, error) {
	data, err := s.source.FromISBN(isbn)
	if err != nil {
		return nil, fmt.Errorf("from source: %v", err)
	}

	params, err := s.process.ProcessSourceData(data)
	if err != nil {
		return nil, fmt.Errorf("process source data: %v", err)
	}

	entity, err := s.repository.Create(*params)
	if err != nil {
		return nil, fmt.Errorf("create: %v", err)
	}

	return entity, nil
}
