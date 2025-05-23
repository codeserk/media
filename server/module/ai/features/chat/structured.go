package chat

import (
	"encoding/json"
	"fmt"
)

func (s *service) Structured(context string, schema any) (any, error) {
	if s.mistral == nil {
		return "", fmt.Errorf("no provider available")
	}

	text, err := s.mistral.Structured(context, schema)
	if err != nil {
		return "", fmt.Errorf("failed to get text: %w", err)
	}

	err = json.Unmarshal([]byte(text), &schema)
	if err != nil {
		return text, fmt.Errorf("failed to unmarshal text into schema: %w", err)
	}

	return text, nil
}
