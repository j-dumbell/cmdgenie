package config

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/openai/openai-go"
)

type Config struct {
	OpenAIAPIKey *string           `json:"openAIApiKey"`
	DefaultModel *openai.ChatModel `json:"defaultModel"`
}

// Service handles saving and loading config from file.
type Service struct {
	filePath string
}

func NewService(filePath string) Service {
	return Service{filePath: filePath}
}

func (service *Service) Load() (Config, error) {
	file, err := os.Open(service.filePath)
	if errors.Is(err, os.ErrNotExist) {
		return Config{}, nil
	}
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func (service *Service) Save(config Config) error {
	file, err := os.Create(service.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(config)
}
