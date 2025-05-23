package config

import (
	"fmt"
)

type Definition string

const (
	DefinitionApiPublic Definition = "api-public"
)

var definitionConfigPath = "./internal/config/definitions/"

var definitionConfigFile = map[Definition]string{
	DefinitionApiPublic: "api-public.yaml",
}

func configFile(definition Definition) (string, error) {
	if config, exists := definitionConfigFile[definition]; exists {
		return config, nil
	}

	return "", fmt.Errorf("invalid definition %s", definition)
}
