package chat

import "fmt"

func (s *service) Text(context string) (string, error) {
	if s.mistral == nil {
		return "", fmt.Errorf("no provider available")
	}

	return s.mistral.Text(context)
}
