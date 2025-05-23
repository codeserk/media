package mongo

import (
	"context"
	"fmt"
	"media/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection struct {
	conf   *config.Mongo
	client *mongo.Client

	Users *mongo.Collection
	Books *mongo.Collection
}

var Users *mongo.Collection

func (c *Connection) Connect(connectionURL string) error {
	cxt, cancel := c.Context()
	defer cancel()

	if connectionURL == "" {
		connectionURL = c.connectionURL()
	}
	clientOptions := options.Client().ApplyURI(connectionURL)
	client, err := mongo.Connect(cxt, clientOptions)
	if err != nil {
		return fmt.Errorf("failed to connect to mongo: %v", err)
	}

	err = client.Ping(cxt, nil)
	if err != nil {
		return fmt.Errorf("failed to ping mongo server: %v", err)
	}
	c.client = client

	// Init collections
	c.createCollections()

	return nil
}

func (c *Connection) createCollections() {
	c.Users = c.client.Database(c.conf.Database).Collection("users")
	c.Books = c.client.Database(c.conf.Database).Collection("books")
}

func (c *Connection) connectionURL() string {
	return fmt.Sprintf("mongodb://%v:%v@%v:%v/%v", c.conf.User, c.conf.Password, c.conf.Host, c.conf.Port, c.conf.Database)
}

func (c *Connection) Context() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

func (c *Connection) Transaction(ctx context.Context, opts ...*options.SessionOptions) (mongo.Session, func(err error), error) {
	session, err := c.client.StartSession(opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("start session: %v", err)
	}
	err = session.StartTransaction()
	if err != nil {
		return nil, nil, fmt.Errorf("start transaction: %v", err)
	}

	finish := func(err error) {
		session.EndSession(ctx)

		if err != nil {
			session.AbortTransaction(ctx)
		} else {
			session.CommitTransaction(ctx)
		}
	}

	return session, finish, err
}
