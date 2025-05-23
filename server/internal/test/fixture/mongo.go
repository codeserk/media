package fixture

import (
	"context"
	"fmt"
	"log"
	"media/internal/mongo"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var TestObjectIdHex = "60b8d295f1b2c3d4e5f6a7b8"
var TestObjectId, _ = primitive.ObjectIDFromHex(TestObjectIdHex)

func Mongo(t *testing.T) *mongo.Connection {
	ctx := context.Background()

	container, err := mongodb.Run(ctx, "mongo:8")
	if err != nil {
		log.Fatalf("create container: %v", err)
	}
	t.Cleanup(func() {
		if err := testcontainers.TerminateContainer(container); err != nil {
			log.Fatalf("terminate container: %v", err)
		}
	})

	connectionURL, err := container.ConnectionString(ctx)
	if err != nil {
		log.Fatalf("get connection url: %v", err)
	}

	connection, err := mongo.FromConnectionString(fmt.Sprintf("%s/press", connectionURL))
	if err != nil {
		log.Fatalf("create connection: %v", connection)
	}

	return connection
}

func ClearMongo(db *mongo.Connection) {
	db.Users.DeleteMany(context.Background(), bson.M{})
}
