package mistral

import (
	"fmt"

	"github.com/gage-technologies/mistral-go"
)

func (s *Service) Structured(context string, schema any) (string, error) {
	res, err := s.client.Chat(s.conf.Model, []mistral.ChatMessage{{
		Content: context, Role: mistral.RoleUser}}, s.structuredParams)
	if err != nil {
		return "", fmt.Errorf("chat: %v", err)
	}
	if len(res.Choices) == 0 {
		return "", fmt.Errorf("no response: %v", err)
	}

	return res.Choices[0].Message.Content, nil
}
