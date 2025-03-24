package commands

import (
	"github.com/Le-BlitzZz/onchain-auth-app/internal/config"
	"github.com/urfave/cli/v2"
)

var InitConfig = func(ctx *cli.Context) *config.Config {
	return config.NewConfig(ctx)
}
