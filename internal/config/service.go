package config

import (
	"encoding/json"
	"os"

	"github.com/openai/openai-go"
)

type Config struct {
	OpenAIAPIKey *string           `json:"openAIApiKey"`
	DefaultModel *openai.ChatModel `json:"defaultModel"`
}

type Service struct {
	filePath string
}

func NewService(filePath string) Service {
	return Service{filePath: filePath}
}

func (service *Service) Load() (*Config, error) {
	file, err := os.Open(service.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (service *Service) Save(config Config) error {
	file, err := os.Create(service.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(config)
}
