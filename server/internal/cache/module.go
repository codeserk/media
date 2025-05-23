package cache

import (
	"context"
	"fmt"
	"media/internal/config"
	"net/url"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type Service struct {
	redis *redis.Client
}

func New(conf *config.Redis) (*Service, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Password,
		DB:       conf.DB,
	})
	status := client.Ping(context.Background())
	if err := status.Err(); err != nil {
		return nil, fmt.Errorf("failed to connect: %v", err)
	}

	return &Service{redis: client}, nil
}

func FromConnectionString(connectionURL string) (*Service, error) {
	parsedURL, err := url.Parse(connectionURL)
	if err != nil {
		return nil, err
	}

	host := parsedURL.Hostname()
	portStr := parsedURL.Port()
	port, err := strconv.ParseUint(portStr, 10, 16)
	if err != nil {
		return nil, fmt.Errorf("database port: %v", err)
	}

	password, _ := parsedURL.User.Password()
	dbStr := parsedURL.Path[1:]
	db, err := strconv.Atoi(dbStr)
	if err != nil {
		return nil, fmt.Errorf("database number: %v", err)
	}

	return New(&config.Redis{
		Host:     host,
		Port:     uint16(port),
		Password: password,
		DB:       db,
	})
}
