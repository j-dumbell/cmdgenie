package cli

import (
	"context"

	"github.com/openai/openai-go"
)

type OpenAIClient interface {
	Ask(ctx context.Context, model openai.ChatModel, promptContext string, prompt string) (string, error)
}

type OpenAIClientFactory func(apiKey string) OpenAIClient
