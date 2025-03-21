package commands

import (
	"context"
	"github.com/Le-BlitzZz/onchain-auth-app/internal/server"
	"github.com/Le-BlitzZz/onchain-auth-app/internal/server/process"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
	"syscall"
)

// StartCommand configures the command name and action.
var StartCommand = &cli.Command{
	Name:   "start",
	Usage:  "Starts the Web server",
	Action: startAction,
}

// startAction starts the Web server.
func startAction(ctx *cli.Context) error {
	cctx, cancel := context.WithCancel(context.Background())

	go server.Start(cctx)

	signal.Notify(process.Signal, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	sig := <-process.Signal

	log.Info("Shutting down...")
	cancel()

	if sig == syscall.SIGTERM {
		os.Exit(1)
		return nil
	}

	return nil
}
