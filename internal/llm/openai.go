package llm

import (
	"context"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type OpenAIClient struct {
	client openai.Client
}

func NewOpenAIClient(apiKey string) OpenAIClient {
	client := openai.NewClient(option.WithAPIKey(apiKey))
	return OpenAIClient{
		client: *client,
	}
}

func (client *OpenAIClient) Ask(ctx context.Context, model openai.ChatModel, promptContext string, prompt string) (string, error) {
	chatCompletion, err := client.client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(promptContext),
			openai.UserMessage(prompt),
		}),
		Model: openai.F(model),
	})
	if err != nil {
		return "", err
	}

	return chatCompletion.Choices[0].Message.Content, nil
}
