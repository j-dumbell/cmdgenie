package cli

import (
	"bytes"
	"context"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/j-dumbell/cmdgenie/internal/config"
	"github.com/j-dumbell/cmdgenie/internal/llm"
	"github.com/j-dumbell/cmdgenie/internal/util"
	"github.com/openai/openai-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestApp_configure(t *testing.T) {
	apiKey := "abc123"
	model := openai.ChatModelGPT4oMini

	configService := newTestConfigService(t)

	app := NewApp(configService, nil, &TestSelectPrompter{model}, &TestTextPrompter{apiKey}, os.Stdout)
	err := app.Run(context.Background(), []string{"cmdgenie", "configure"})
	require.NoError(t, err, "Run should not error")

	actual, err := configService.Load()
	require.NoError(t, err, "failed to load config")

	expected := config.Config{
		OpenAIAPIKey: &apiKey,
		DefaultModel: &model,
	}
	assert.Equal(t, expected, actual, "saved config")
}

func TestApp_ask(t *testing.T) {
	configService := newTestConfigService(t)
	savedConfig := config.Config{
		OpenAIAPIKey: util.ToPtr("abc123"),
		DefaultModel: util.ToPtr(openai.ChatModelGPT4oMini),
	}
	err := configService.Save(savedConfig)
	require.NoError(t, err, "failed to save config")

	response := "ls -a"
	openAIClientFactory, calls := NewTestOpenAIClientFactory(response)
	ask := "list all files"

	outputBuffer := bytes.Buffer{}
	app := NewApp(configService, openAIClientFactory, &TestSelectPrompter{""}, &TestTextPrompter{""}, &outputBuffer)

	err = app.Run(context.Background(), []string{"cmdgenie", "ask", ask})
	require.NoError(t, err, "Run should not error")

	require.Equal(t, 1, len(*calls), "expected 1 OpenAI API call")
	call := (*calls)[0]
	assert.Equal(t, *savedConfig.DefaultModel, call.Model)
	assert.Equal(t, ask, call.Prompt)

	assert.Equal(t, response+"\n", outputBuffer.String(), "output")
}

func TestApp_list_models(t *testing.T) {
	configService := newTestConfigService(t)
	openAIClientFactory, _ := NewTestOpenAIClientFactory("")
	outputBuffer := bytes.Buffer{}
	app := NewApp(configService, openAIClientFactory, &TestSelectPrompter{""}, &TestTextPrompter{""}, &outputBuffer)

	err := app.Run(context.Background(), []string{"cmdgenie", "list-models"})
	require.NoError(t, err, "Run should not error")

	expected := strings.Join(llm.Models, "\n") + "\n"
	assert.Equal(t, expected, outputBuffer.String(), "output")
}

func newTestConfigService(t *testing.T) config.Service {
	dir := t.TempDir()
	filePath := path.Join(dir, ".cmdgenie.json")
	return config.NewService(filePath)
}

type Call struct {
	Model         openai.ChatModel
	PromptContext string
	Prompt        string
}

type TestOpenAIClient struct {
	Calls    *[]Call
	Response string
}

func (client *TestOpenAIClient) Ask(_ context.Context, model openai.ChatModel, promptContext string, prompt string) (string, error) {
	newCalls := append(*client.Calls, Call{
		Model:         model,
		PromptContext: promptContext,
		Prompt:        prompt,
	})
	*client.Calls = newCalls
	return client.Response, nil
}

func NewTestOpenAIClientFactory(response string) (OpenAIClientFactory, *[]Call) {
	calls := []Call{}
	return func(apiKey string) OpenAIClient {
		return &TestOpenAIClient{
			Response: response,
			Calls:    &calls,
		}
	}, &calls
}

type TestTextPrompter struct {
	Text string
}

func (prompter *TestTextPrompter) Run() (string, error) {
	return prompter.Text, nil
}

type TestSelectPrompter struct {
	Selected string
}

func (prompter TestSelectPrompter) Run() (int, string, error) {
	return 0, prompter.Selected, nil
}
