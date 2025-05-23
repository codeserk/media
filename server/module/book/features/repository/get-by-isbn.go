package repository

import "media/module/book"

func (r *repository) GetByISBN(isbn book.ISBN) (*entity, error) {
	ctx, cancel := r.db.Context()
	defer cancel()

	return r.base.GetOne(ctx, book.FilterParams{ISBN: string(isbn)})
}
