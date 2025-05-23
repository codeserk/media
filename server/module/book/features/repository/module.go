package repository

import (
	"media/internal/mongo"
)

type repository struct {
	db   *mongo.Connection
	base *mongo.Repository[entity, filterParams, createParams, updateParams, document]
}

func New(db *mongo.Connection) *repository {
	return &repository{
		db:   db,
		base: mongo.NewRepository(db.Books, toFilterQuery, toCreateQuery, toUpdateQuery, toEntity),
	}
}
