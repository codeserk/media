package config

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

const definitionsFolder = "./internal/definitions/"

var environmentConfigFile = map[string]string{
	"development": "development.yaml",
	"staging":     "staging.yaml",
	"production":  "production.yaml",
}

var validate *validator.Validate

func Parse[TConfig any](definition Definition, conf *TConfig) error {
	file, err := configFile(definition)
	if err != nil {
		return fmt.Errorf("get config file: %v", err)
	}

	// Read config from file
	viper.AddConfigPath(definitionConfigPath)
	viper.SetConfigFile(definitionConfigPath + file)
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("read file: %v", err)
	}

	validate = validator.New()

	// Read config from Env
	err = godotenv.Load()
	if err != nil {
		log.Info().Msgf("no .env: %v", err)
	}

	viper.SetEnvPrefix("app")
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
	viper.AutomaticEnv()

	err = viper.Unmarshal(&conf)
	if err != nil {
		return fmt.Errorf("unmarshal config: %v", err)
	}
	err = validate.Struct(conf)
	if err != nil {
		return fmt.Errorf("validate config: %v", err)
	}

	return nil
}
