package main

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/j-dumbell/cmdgenie/internal/cli"
	"github.com/j-dumbell/cmdgenie/internal/config"
)

var configFileName = ".cmdgenie.json"

func main() {
	homeDir, _ := os.UserHomeDir()
	configFilePath := path.Join(homeDir, configFileName)

	configService := config.NewService(configFilePath)
	app := cli.NewApp(configService)

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
}
