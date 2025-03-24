package main

import (
	"github.com/Le-BlitzZz/onchain-auth-app/internal/commands"
	"github.com/Le-BlitzZz/onchain-auth-app/internal/config"
	"github.com/Le-BlitzZz/onchain-auth-app/internal/event"
	"github.com/urfave/cli/v2"
	"os"
)

var version = "development"
var log = event.Log

const appName = "AuthOnchain"
const appAbout = "AuthOnchain"

// Metadata contains build specific information.
var Metadata = map[string]interface{}{
	"Name":    appName,
	"About":   appAbout,
	"Version": version,
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			os.Exit(1)
		}
	}()

	app := cli.NewApp()
	app.Usage = appAbout
	app.Version = version
	app.EnableBashCompletion = true
	app.Flags = config.Flags.Cli()
	app.Commands = commands.AuthOnchain
	app.Metadata = Metadata

	if err := app.Run(os.Args); err != nil {
		log.Error(err)
	}
}
