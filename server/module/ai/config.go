package ai

type Provider = string

var (
	ProviderMistral Provider = "mistral"
)

type MistralConfig struct {
	ApiKey      string  `json:"apiKey" validate:"required"`
	Model       string  `json:"model" validate:"required,oneof=mistral-large-latest mistral-medium-latest mistral-tiny"`
	Temperature float64 `json:"temperature"`
	TopP        float64 `json:"topP"`
	MaxTokens   int     `json:"maxTokens"`
}

type Config struct {
	Provider Provider       `json:"provider" validate:"required,oneof=mistral"`
	Mistral  *MistralConfig `json:"mistral"`
}
