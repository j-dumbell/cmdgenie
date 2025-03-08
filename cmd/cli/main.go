package main

import (
	"context"
	"fmt"
	"os"

	"github.com/j-dumbell/cmdgenie/internal/cli"
	"github.com/j-dumbell/cmdgenie/internal/config"
)

const configFilePath = "~/.cmdgenie.json"

func main() {
	configService := config.NewService(configFilePath)
	app := cli.NewApp(configService)

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
}
