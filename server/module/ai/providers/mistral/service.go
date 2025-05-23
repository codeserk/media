package mistral

import (
	"media/module/ai"

	"github.com/gage-technologies/mistral-go"
)

type Service struct {
	conf             *ai.MistralConfig
	client           *mistral.MistralClient
	params           *mistral.ChatRequestParams
	structuredParams *mistral.ChatRequestParams
}

func New(conf *ai.MistralConfig) *Service {
	client := mistral.NewMistralClientDefault(conf.ApiKey)
	params := mistral.ChatRequestParams{
		Temperature: 0.5,
		TopP:        0.8,
		RandomSeed:  42069,
		MaxTokens:   500,
		SafePrompt:  false,
	}
	structuredParams := mistral.ChatRequestParams(params)
	structuredParams.ResponseFormat = mistral.ResponseFormatJsonObject

	if conf.Temperature != 0 {
		params.Temperature = float64(conf.Temperature)
	}
	if conf.TopP != 0 {
		params.TopP = float64(conf.TopP)
	}
	if conf.MaxTokens != 0 {
		params.MaxTokens = int(conf.MaxTokens)
	}

	return &Service{conf, client, &params, &structuredParams}
}
