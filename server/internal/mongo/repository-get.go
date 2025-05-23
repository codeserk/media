package mongo

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// entity

func (r *Repository[TEntity, TFilterParams, TCreateParams, TUpdateParams, TDocument]) GetById(ctx context.Context, id string) (*TEntity, error) {
	d, err := r.GetDocumentById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get document: %v", err)
	}
	if d == nil {
		return nil, nil
	}

	return r.toEntity(d), nil
}

func (r *Repository[TEntity, TFilterParams, TCreateParams, TUpdateParams, TDocument]) GetOne(ctx context.Context, params TFilterParams) (*TEntity, error) {
	d, err := r.GetOneDocument(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("get document: %v", err)
	}
	if d == nil {
		return nil, nil
	}

	return r.toEntity(d), nil
}

func (r *Repository[TEntity, TFilterParams, TCreateParams, TUpdateParams, TDocument]) GetOneUsingQuery(ctx context.Context, query any) (*TEntity, error) {
	d, err := r.GetOneDocumentUsingQuery(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("get document: %v", err)
	}
	if d == nil {
		return nil, nil
	}

	return r.toEntity(d), nil
}

func (r *Repository[TEntity, TFilterParams, TCreateParams, TUpdateParams, TDocument]) GetMany(ctx context.Context, params TFilterParams) ([]*TEntity, error) {
	docs, err := r.GetManyDocuments(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("get documents: %v", err)
	}

	entities := make([]*TEntity, len(docs))
	for i, doc := range docs {
		entities[i] = r.toEntity(doc)
	}

	return entities, nil
}

func (r *Repository[TEntity, TFilterParams, TCreateParams, TUpdateParams, TDocument]) GetManyUsingQuery(ctx context.Context, query any) ([]*TEntity, error) {
	docs, err := r.GetManyDocumentsUsingQuery(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("get documents: %v", err)
	}

	entities := make([]*TEntity, len(docs))
	for i, doc := range docs {
		entities[i] = r.toEntity(doc)
	}

	return entities, nil
}

// document

func (r *Repository[TEntity, TFilterParams, TCreateParams, TUpdateParams, TDocument]) GetDocumentById(ctx context.Context, id string) (*TDocument, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("transform id: %v", err)
	}

	var document TDocument
	err = r.collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&document)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, fmt.Errorf("find one: %v", err)
	}

	return &document, nil
}

func (r *Repository[TEntity, TFilterParams, TCreateParams, TUpdateParams, TDocument]) GetOneDocument(ctx context.Context, params TFilterParams) (*TDocument, error) {
	query, err := r.toFilterQuery(params)
	if err != nil {
		return nil, fmt.Errorf("filter query: %v", err)
	}

	return r.GetOneDocumentUsingQuery(ctx, query)
}

func (r *Repository[TEntity, TFilterParams, TCreateParams, TUpdateParams, TDocument]) GetOneDocumentUsingQuery(ctx context.Context, query any) (*TDocument, error) {
	var document TDocument
	err := r.collection.FindOne(ctx, query).Decode(&document)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, fmt.Errorf("find one: %v", err)
	}

	return &document, nil
}

func (r *Repository[TEntity, TFilterParams, TCreateParams, TUpdateParams, TDocument]) GetManyDocuments(ctx context.Context, params TFilterParams) ([]*TDocument, error) {
	query, err := r.toFilterQuery(params)
	if err != nil {
		return nil, fmt.Errorf("filter query: %v", err)
	}

	return r.GetManyDocumentsUsingQuery(ctx, query)
}

func (r *Repository[TEntity, TFilterParams, TCreateParams, TUpdateParams, TDocument]) GetManyDocumentsUsingQuery(ctx context.Context, query any) ([]*TDocument, error) {
	cursor, err := r.collection.Find(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("find: %v", err)
	}
	defer cursor.Close(ctx)

	var documents []*TDocument
	for cursor.Next(ctx) {
		var document TDocument
		if err := cursor.Decode(&document); err != nil {
			return nil, fmt.Errorf("decode: %v", err)
		}
		documents = append(documents, &document)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return documents, nil
}
