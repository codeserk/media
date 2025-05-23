package ai

type ChatService interface {
	Text(context string) (string, error)
	Structured(context string, schema any) (any, error)
}
