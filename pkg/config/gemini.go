package config

import (
	"errors"
	"os"

	"github.com/samber/do/v2"
)

type GeminiConfig struct {
	APIKey string
}

func NewGeminiConfig(_ do.Injector) (*GeminiConfig, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return nil, errors.New("GEMINI_API_KEY is not set")
	}

	return &GeminiConfig{
		APIKey: apiKey,
	}, nil
}
