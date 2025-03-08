package llm

import (
	"context"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type OpenAIModel struct {
	client *openai.Client
	name   openai.ChatModel
}

func NewOpenAIModel(apiKey string, modelName openai.ChatModel) OpenAIModel {
	client := openai.NewClient(option.WithAPIKey(apiKey))
	return OpenAIModel{
		client: client,
		name:   modelName,
	}
}

func (model *OpenAIModel) Ask(ctx context.Context, promptContext string, prompt string) (string, error) {
	chatCompletion, err := model.client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(promptContext),
			openai.UserMessage(prompt),
		}),
		Model: openai.F(model.name),
	})
	if err != nil {
		return "", err
	}

	return chatCompletion.Choices[0].Message.Content, nil
}
