package repository

import (
	"media/module/book"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Types

type entity = book.Entity
type filterParams = book.FilterParams
type createParams = book.CreateParams
type updateParams = book.UpdateParams

type document struct {
	Id       string `bson:"_id"`
	Metadata book.Metadata
	Sources  book.SourceMultiData
	Images   book.Images

	CreatedAt  time.Time    `bson:"createdAt"`
	UpdatedAt  time.Time    `bson:"updatedAt"`
	ArchivedAt *time.Ticker `bson:"archivedAt"`
}

// Queries

type createQuery struct {
	Metadata book.Metadata
	Sources  book.SourceMultiData
	Images   book.Images

	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

type updateQuery struct {
	Metadata book.Metadata
	Sources  book.SourceMultiData
	Images   book.Images

	UpdatedAt time.Time `bson:"updatedAt"`
}

// Transformers

func toFilterQuery(filters filterParams) (any, error) {
	query := bson.M{}
	if filters.Title != "" {
		query["title"] = bson.M{"$regex": filters.Title, "$options": "i"}
	}
	if filters.TitleExact != "" {
		query["title"] = filters.TitleExact
	}
	if filters.ISBN != "" {
		query["metadata.isbn"] = filters.ISBN
	}

	return query, nil
}

func toCreateQuery(params createParams) (any, error) {
	return createQuery{
		Metadata: params.Metadata,
		Sources:  params.Sources,
		Images:   params.Images,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func toUpdateQuery(params updateParams) (any, error) {
	return updateQuery{
		Metadata: params.Metadata,
		Images:   params.Images,

		UpdatedAt: time.Now(),
	}, nil
}

func toEntity(d *document) *entity {
	return &entity{
		Id:       d.Id,
		Metadata: d.Metadata,
		Sources:  d.Sources,
		Images:   d.Images,

		CreatedAt:  d.CreatedAt,
		UpdatedAt:  d.UpdatedAt,
		ArchivedAt: d.ArchivedAt,
	}
}

// Methods

func (s *repository) GetById(id string) (*entity, error) {
	ctx, cancel := s.db.Context()
	defer cancel()

	return s.base.GetById(ctx, id)
}
func (s *repository) GetOne(filters filterParams) (*entity, error) {
	ctx, cancel := s.db.Context()
	defer cancel()

	return s.base.GetOne(ctx, filters)
}
func (s *repository) GetMany(filters filterParams) ([]*entity, error) {
	ctx, cancel := s.db.Context()
	defer cancel()

	return s.base.GetMany(ctx, filters)
}

func (s *repository) Create(params createParams) (*entity, error) {
	ctx, cancel := s.db.Context()
	defer cancel()

	return s.base.Create(ctx, params)
}

func (s *repository) Update(id string, params updateParams) (*entity, error) {
	ctx, cancel := s.db.Context()
	defer cancel()

	return s.base.UpdateById(ctx, id, params)
}

func (s *repository) Archive(id string) (*entity, error) {
	ctx, cancel := s.db.Context()
	defer cancel()

	return s.base.Archive(ctx, id)
}
func (s *repository) Delete(id string) error {
	ctx, cancel := s.db.Context()
	defer cancel()

	return s.base.Delete(ctx, id)
}
