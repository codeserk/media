package main

import (
	"fmt"
	"media/internal/config"
	"media/internal/email"
	"media/internal/media"
	"media/module/ai"
	"media/module/book"
)

type Config struct {
	API       config.API       `json:"api" validate:"required"`
	Mongo     config.Mongo     `json:"mongo" validate:"required"`
	Redis     config.Redis     `json:"redis" validate:"required"`
	JWT       config.JWT       `json:"jwt" validate:"required"`
	Email     email.Config     `json:"email" validate:"required"`
	Dashboard config.Dashboard `json:"dashboard" validate:"required"`
	Media     media.Config     `json:"media" validate:"required"`

	BookSource book.SourceConfig `json:"bookSource" validate:"required"`
	AI         ai.Config         `json:"ai" validate:"required"`
}

func ParseConfig() (*Config, error) {
	var conf Config
	err := config.Parse(config.DefinitionApiPublic, &conf)
	if err != nil {
		return nil, fmt.Errorf("parse: %v", err)
	}

	return &conf, nil
}
