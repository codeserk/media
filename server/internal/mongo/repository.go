package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository[
	TEntity any,
	TFilterParams any,
	TCreateParams any,
	TUpdateParams any,
	TDocument any,
] struct {
	collection *mongo.Collection

	toFilterQuery func(params TFilterParams) (any, error)
	toCreateQuery func(params TCreateParams) (any, error)
	toUpdateQuery func(params TUpdateParams) (any, error)
	toEntity      func(document *TDocument) *TEntity
}

func NewRepository[
	TEntity any,
	TFilterParams any,
	TCreateParams any,
	TUpdateParams any,
	TDocument any,
](
	collection *mongo.Collection,
	toFilterQuery func(params TFilterParams) (any, error),
	toCreateQuery func(params TCreateParams) (any, error),
	toUpdateQuery func(params TUpdateParams) (any, error),
	toEntity func(*TDocument,
	) *TEntity) *Repository[TEntity, TFilterParams, TCreateParams, TUpdateParams, TDocument] {
	return &Repository[TEntity, TFilterParams, TCreateParams, TUpdateParams, TDocument]{
		collection,

		toFilterQuery,
		toCreateQuery,
		toUpdateQuery,
		toEntity,
	}
}

// Create

func (r *Repository[TEntity, TFilterParams, TCreateParams, TUpdateParams, TDocument]) Create(ctx context.Context, params TCreateParams) (*TEntity, error) {

	d, err := r.CreateDocument(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("create document: %v", err)
	}

	return r.toEntity(d), nil
}

func (r *Repository[TEntity, TFilterParams, TCreateParams, TUpdateParams, TDocument]) CreateDocument(ctx context.Context, params TCreateParams) (*TDocument, error) {
	query, err := r.toCreateQuery(params)
	if err != nil {
		return nil, fmt.Errorf("create query: %v", err)
	}

	result, err := r.collection.InsertOne(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("insert: %v", err)
	}
	if objectId, ok := result.InsertedID.(primitive.ObjectID); ok {
		return r.GetDocumentById(ctx, objectId.Hex())
	}

	return nil, fmt.Errorf("invalid inserted id")
}

// Delete

func (r *Repository[TEntity, TFilterParams, TCreateParams, TUpdateParams, TDocument]) Delete(ctx context.Context, id string) error {
	return r.DeleteDocument(ctx, id)
}

func (r *Repository[TEntity, TFilterParams, TCreateParams, TUpdateParams, TDocument]) DeleteDocument(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("transform id: %v", err)
	}
	filter := bson.M{"_id": objectId}
	_, err = r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("delete: %v", err)
	}

	return nil
}

// Archive

func (r *Repository[TEntity, TFilterParams, TCreateParams, TUpdateParams, TDocument]) Archive(ctx context.Context, id string) (*TEntity, error) {
	return r.UpdateByIdUsingQuery(ctx, id, bson.M{
		"$set": bson.M{"archivedAt": time.Now()},
	})
}

func (r *Repository[TEntity, TFilterParams, TCreateParams, TUpdateParams, TDocument]) ArchiveDocument(ctx context.Context, id string) (*TDocument, error) {
	return r.UpdateDocumentByIdUsingQuery(ctx, id, bson.M{
		"$set": bson.M{"archivedAt": time.Now()},
	})
}
