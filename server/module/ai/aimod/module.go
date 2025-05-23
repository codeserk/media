package aimod

import (
	"media/module/ai"
	"media/module/ai/features/chat"
)

func NewChatService(conf *ai.Config) ai.ChatService {
	return chat.New(conf)
}
