package commands

import (
	"github.com/Le-BlitzZz/onchain-auth-app/internal/event"
	"github.com/urfave/cli/v2"
)

var log = event.Log

var AuthOnchain = []*cli.Command{
	StartCommand,
}
