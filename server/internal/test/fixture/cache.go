package fixture

import (
	"context"
	"fmt"
	"log"
	"media/internal/cache"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/redis"
)

func Cache(t *testing.T) *cache.Service {
	ctx := context.Background()

	container, err := redis.Run(ctx, "redis:7-alpine")
	if err != nil {
		log.Fatalf("create container: %v", err)
	}
	t.Cleanup(func() {
		if err := testcontainers.TerminateContainer(container); err != nil {
			log.Fatalf("terminate container: %v", err)
		}
	})

	connectionUrl, err := container.ConnectionString(ctx)
	if err != nil {
		log.Fatalf("redis connection url: %v", err)
	}

	service, err := cache.FromConnectionString(fmt.Sprintf("%s/0", connectionUrl))
	if err != nil {
		log.Fatalf("create redis: %v", err)
	}

	return service
}

func ClearCache(cache *cache.Service) {
	cache.Clear()
}
