package main

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/j-dumbell/cmdgenie/internal/cli"
	"github.com/j-dumbell/cmdgenie/internal/config"
	"github.com/j-dumbell/cmdgenie/internal/llm"
)

var configFileName = ".cmdgenie.json"

func main() {
	homeDir, _ := os.UserHomeDir()
	configFilePath := path.Join(homeDir, configFileName)
	configService := config.NewService(configFilePath)

	openAIFactory := func(apiKey string) cli.OpenAIClient {
		client := llm.NewOpenAIClient(apiKey)
		return &client
	}

	app := cli.NewApp(configService, openAIFactory, &cli.ModelSelect, &cli.ApiKeyPrompt, os.Stdout)

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
}
