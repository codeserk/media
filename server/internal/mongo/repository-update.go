package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// entity

func (r *Repository[TEntity, TFilterParams, TCreateParams, TUpdateParams, TDocument]) UpdateById(ctx context.Context, id string, params TUpdateParams) (*TEntity, error) {
	d, err := r.UpdateDocumentById(ctx, id, params)
	if err != nil {
		return nil, fmt.Errorf("update document: %v", err)
	}
	if d == nil {
		return nil, nil
	}

	return r.toEntity(d), nil
}

func (r *Repository[TEntity, TFilterParams, TCreateParams, TUpdateParams, TDocument]) UpdateByIdUsingQuery(ctx context.Context, id string, query any, opts ...*options.UpdateOptions) (*TEntity, error) {
	d, err := r.UpdateDocumentByIdUsingQuery(ctx, id, query, opts...)
	if err != nil {
		return nil, fmt.Errorf("update document: %v", err)
	}
	if d == nil {
		return nil, nil
	}

	return r.toEntity(d), nil
}

// document

func (r *Repository[TEntity, TFilterParams, TCreateParams, TUpdateParams, TDocument]) UpdateOneUsingQuery(ctx context.Context, filterQuery any, updateQuery any, opts ...*options.UpdateOptions) (*TEntity, error) {
	d, err := r.UpdateOneDocumentUsingQuery(ctx, filterQuery, updateQuery, opts...)
	if err != nil {
		return nil, fmt.Errorf("get one document using query: %v", err)
	}
	if d == nil {
		return nil, nil
	}

	return r.toEntity(d), nil
}

func (r *Repository[TEntity, TFilterParams, TCreateParams, TUpdateParams, TDocument]) UpdateDocumentById(ctx context.Context, id string, params TUpdateParams) (*TDocument, error) {
	query, err := r.toUpdateQuery(params)
	if err != nil {
		return nil, fmt.Errorf("update query: %v", err)
	}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("transform id: %v", err)
	}

	result, err := r.collection.UpdateByID(ctx, objectId, query)
	if err != nil {
		return nil, fmt.Errorf("update: %v", err)
	}
	if objectId, ok := result.UpsertedID.(primitive.ObjectID); ok {
		return r.GetDocumentById(ctx, objectId.Hex())
	}

	return r.GetDocumentById(ctx, id)
}

func (r *Repository[TEntity, TFilterParams, TCreateParams, TUpdateParams, TDocument]) UpdateDocumentByIdUsingQuery(ctx context.Context, id string, query any, opts ...*options.UpdateOptions) (*TDocument, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("transform id: %v", err)
	}
	result, err := r.collection.UpdateByID(ctx, objectId, query, opts...)
	if err != nil {
		return nil, fmt.Errorf("update: %v", err)
	}
	if objectId, ok := result.UpsertedID.(primitive.ObjectID); ok {
		return r.GetDocumentById(ctx, objectId.Hex())
	}

	return r.GetDocumentById(ctx, id)
}

func (r *Repository[TEntity, TFilterParams, TCreateParams, TUpdateParams, TDocument]) UpdateOneDocumentUsingQuery(ctx context.Context, filterQuery, updateQuery any, opts ...*options.UpdateOptions) (*TDocument, error) {
	result, err := r.collection.UpdateOne(ctx, filterQuery, updateQuery, opts...)
	if err != nil {
		return nil, fmt.Errorf("update: %v", err)
	}
	if objectId, ok := result.UpsertedID.(primitive.ObjectID); ok {
		return r.GetDocumentById(ctx, objectId.Hex())
	}

	return r.GetOneDocumentUsingQuery(ctx, filterQuery)
}
