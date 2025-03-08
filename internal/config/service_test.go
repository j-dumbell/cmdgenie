package config

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/openai/openai-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_Save(t *testing.T) {
	dir := t.TempDir()
	filePath := fmt.Sprintf("%s/file.json", dir)
	service := NewService(filePath)

	config := Config{
		OpenAIAPIKey: ToPtr("abc123"),
		DefaultModel: ToPtr(openai.ChatModelGPT4oMini),
	}

	err := service.Save(config)
	require.NoError(t, err, "Save should not error")

	file, err := os.Open(filePath)
	require.NoError(t, err, "failed to open file")
	defer file.Close()

	var savedConfig Config
	err = json.NewDecoder(file).Decode(&savedConfig)
	require.NoError(t, err, "failed to decode file contents")
	assert.Equal(t, config, savedConfig)
}

func TestService_Load(t *testing.T) {
	dir := t.TempDir()
	filePath := fmt.Sprintf("%s/file.json", dir)
	service := NewService(filePath)

	file, err := os.Create(filePath)
	require.NoError(t, err, "failed to create config file")
	defer file.Close()

	configJSON := `{"openAIApiKey":"abc","defaultModel":"gpt-4o-mini"}`
	_, err = file.WriteString(configJSON)
	require.NoError(t, err, "failed to write to config file")

	actual, err := service.Load()
	require.NoError(t, err, "Load should not error")

	expected := Config{
		OpenAIAPIKey: ToPtr("abc"),
		DefaultModel: ToPtr(openai.ChatModelGPT4oMini),
	}

	assert.Equal(t, expected, actual)
}

func TestService_Load_fileNotExists(t *testing.T) {
	dir := t.TempDir()
	filePath := fmt.Sprintf("%s/file.json", dir)
	service := NewService(filePath)

	actual, err := service.Load()
	require.NoError(t, err, "Load should not error")

	expected := Config{
		OpenAIAPIKey: nil,
		DefaultModel: nil,
	}

	assert.Equal(t, expected, actual)
}

func ToPtr[T any](t T) *T {
	return &t
}
