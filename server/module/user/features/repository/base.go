package repository

import (
	"media/module/user"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Types

type entity = user.Entity
type filterParams = user.FilterParams
type createParams = user.CreateParams
type updateParams = user.UpdateParams

type emailVerifyDocument struct {
	IsVerified bool   `bson:"isVerified"`
	Token      string `bson:"token"`
}

type document struct {
	Id            string              `bson:"_id"`
	Name          string              `bson:"name"`
	Email         string              `bson:"email"`
	EmailVerified emailVerifyDocument `bson:"emailVerify"`
	Password      string              `bson:"password"`

	Role user.Role `bson:"role"`

	CreatedAt  time.Time    `bson:"createdAt"`
	UpdatedAt  time.Time    `bson:"updatedAt"`
	ArchivedAt *time.Ticker `bson:"archivedAt"`
}

// Queries

type createQuery struct {
	Name        string              `bson:"name"`
	Email       string              `bson:"email"`
	EmailVerify emailVerifyDocument `bson:"emailVerify"`
	Password    string              `bson:"password"`

	Role user.Role

	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

type updateQuery struct {
	Name string `bson:"name"`

	UpdatedAt time.Time `bson:"updatedAt"`
}

type permissionsQuery struct {
	Read   bool `bson:"read"`
	Update bool `bson:"update"`
	Delete bool `bson:"delete"`
}

// Transformers

func toFilterQuery(filters filterParams) (any, error) {
	query := bson.M{}
	if filters.Name != "" {
		query["name"] = bson.M{"$regex": filters.Name, "$options": "i"}
	}
	if filters.NameExact != "" {
		query["name"] = filters.NameExact
	}

	if filters.Email != "" {
		query["email"] = bson.M{"$regex": filters.Email, "$options": "i"}
	}
	if filters.EmailExact != "" {
		query["email"] = filters.EmailExact
	}

	return query, nil
}

func toCreateQuery(params createParams) (any, error) {
	return createQuery{
		Name:        params.Name,
		Email:       params.Email,
		EmailVerify: emailVerifyDocument(params.EmailVerify),
		Password:    params.Password,

		Role: params.Role,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func toUpdateQuery(params updateParams) (any, error) {
	return updateQuery{
		Name: params.Name,

		UpdatedAt: time.Now(),
	}, nil
}

func toEntity(d *document) *entity {
	return &entity{
		Id:          d.Id,
		Name:        d.Name,
		Email:       d.Email,
		EmailVerify: user.EmailVerify(d.EmailVerified),
		Password:    d.Password,

		Role: d.Role,

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
