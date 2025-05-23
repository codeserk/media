package chat

import (
	"media/module/ai"
	"media/module/ai/providers/mistral"
)

type service struct {
	conf    *ai.Config
	mistral *mistral.Service
}

func New(conf *ai.Config) ai.ChatService {
	s := service{conf, nil}
	if conf.Mistral != nil {
		s.mistral = mistral.New(conf.Mistral)
	}

	return &s
}
