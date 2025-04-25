package commands

import (
	"github.com/Le-BlitzZz/onchain-auth-app/internal/config"
	"github.com/urfave/cli/v2"
)

func InitConfig(ctx *cli.Context) (*config.Config, error) {
	cfg := config.NewConfig(ctx)
	return cfg, cfg.Init()
}
