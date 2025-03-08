package cli

import (
	"context"
	"errors"
	"fmt"

	"github.com/j-dumbell/cmdgenie/internal/chatcontext"
	"github.com/j-dumbell/cmdgenie/internal/config"
	"github.com/j-dumbell/cmdgenie/internal/llm"
	"github.com/manifoldco/promptui"
	"github.com/openai/openai-go"
	"github.com/urfave/cli/v3"
)

var (
	ErrMissingPrompt           = errors.New("please provide a prompt")
	ErrMultiplePromptsProvided = errors.New("only a single prompt may be provided")
	ErrOpenAIAPIKeyNotSet      = errors.New("no OpenAI API key found")
)

var (
	modelFlag = &cli.StringFlag{
		Name:     "model",
		Value:    openai.ChatModelGPT4oMini,
		Usage:    "Specify OpenAI model",
		Aliases:  []string{"m"},
		Required: false,
	}

	verboseFlag = &cli.BoolFlag{
		Name:     "verbose",
		Usage:    "whether to return longer responses with explanations and examples",
		Aliases:  []string{"v"},
		Required: false,
	}
)

func NewApp(configService config.Service) cli.Command {
	return cli.Command{
		Name:  "cmdgenie",
		Usage: "Generate shell commands using AI",
		Commands: []*cli.Command{
			{
				Name:  "ask",
				Usage: "Execute a prompt and return a command",
				Flags: []cli.Flag{modelFlag, verboseFlag},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					prompt := cmd.Args().Get(0)
					if prompt == "" {
						return ErrMissingPrompt
					}

					if cmd.Args().Len() > 1 {
						return ErrMultiplePromptsProvided
					}

					cfg, err := configService.Load()
					if err != nil {
						return err
					}

					if cfg.OpenAIAPIKey == nil {
						return ErrOpenAIAPIKeyNotSet
					}

					promptContext := chatcontext.Minimal
					if cmd.Bool(verboseFlag.Name) {
						promptContext = chatcontext.Verbose
					}

					openAIClient := llm.NewOpenAIClient(*cfg.OpenAIAPIKey)
					response, err := openAIClient.Ask(ctx, cmd.String(modelFlag.Name), promptContext, prompt)
					if err != nil {
						return err
					}

					fmt.Println(response)
					return nil
				},
			},
			{
				Name:  "configure",
				Usage: "Configure OpenAI API key and default model",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					modelSelect := promptui.Select{
						Label: "Select model",
						Items: llm.Models,
						Size:  len(llm.Models),
					}

					_, model, err := modelSelect.Run()
					if err != nil {
						return err
					}

					apiKeyPrompt := promptui.Prompt{
						Label: "Enter OpenAI API key",
						Mask:  '*',
					}

					apiKey, err := apiKeyPrompt.Run()
					if err != nil {
						return err
					}

					config := config.Config{
						OpenAIAPIKey: &apiKey,
						DefaultModel: &model,
					}
					return configService.Save(config)
				},
			},
		},
	}
}
