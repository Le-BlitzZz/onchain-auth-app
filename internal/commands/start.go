package commands

import (
	"context"
	"github.com/Le-BlitzZz/onchain-auth-app/internal/server"
	"github.com/Le-BlitzZz/onchain-auth-app/internal/server/process"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// StartCommand configures the command name and action.
var StartCommand = &cli.Command{
	Name:   "start",
	Usage:  "Starts the Web server",
	Action: startAction,
}

// startAction starts the Web server.
func startAction(ctx *cli.Context) error {
	conf := InitConfig(ctx)

	if conf.HttpPort() < 1 || conf.HttpPort() > 65535 {
		log.Fatal("Server port must be a number between 1 and 65535")
	}

	// Pass this context down the chain.
	cctx, cancel := context.WithCancel(context.Background())

	// Start built-in web server.
	go server.Start(cctx, conf)

	// Wait for signal to trigger server shutdown or restart.
	signal.Notify(process.Signal, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-process.Signal

	log.Info("Shutting down...")
	cancel()

	time.Sleep(2 * time.Second)

	return nil
}
