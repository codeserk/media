package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

func (s *Service) String(key string) (string, error) {
	value, err := s.redis.Get(context.Background(), key).Result()
	if err != nil {
		return "", fmt.Errorf("failed to get key: %v", err)
	}

	return value, nil
}
func (s *Service) Struct(key string, out any) error {
	value, err := s.String(key)
	if err != nil {
		return fmt.Errorf("failed to get string value: %v", err)
	}

	if err := json.Unmarshal([]byte(value), out); err != nil {
		return fmt.Errorf("failed to unmarshal value: %v", err)
	}

	return nil
}

func (s *Service) SetString(key string, value string, ttl time.Duration) error {
	status := s.redis.Set(context.Background(), key, value, ttl)
	if err := status.Err(); err != nil {
		return fmt.Errorf("failed to set key: %v", err)
	}

	return nil
}

func (s *Service) SetStruct(key string, value any, ttl time.Duration) error {
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal value: %v", err)
	}

	if err := s.SetString(key, string(jsonBytes), ttl); err != nil {
		return fmt.Errorf("failed to set string value: %v", err)
	}

	return nil
}

func (s *Service) Clear() {
	err := s.redis.FlushDB(context.Background()).Err()
	if err != nil {
		fmt.Printf("failed to clear redis: %v\n", err)
	}
}
