package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/Le-BlitzZz/onchain-auth-app/internal/config"
	"github.com/Le-BlitzZz/onchain-auth-app/internal/server/process"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Start the REST API server.
func Start(ctx context.Context, conf *config.Config) {
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
		}
	}()

	start := time.Now()

	// Log the server process ID for troubleshooting purposes.
	log.Infof("Server: started as pid %d", process.ID)

	// Create default router engine.
	router := gin.Default()

	// Create a new HTTP server instance.
	server := &http.Server{
		Handler: router,
	}

	tcpSocket := fmt.Sprintf("%s:%d", conf.HttpHost(), conf.HttpPort())

	server.Addr = tcpSocket

	log.Infof("Server: listening on %s [%s]", server.Addr, time.Since(start))

	// Start Web server.
	go StartHttp(server)

	// Graceful web server shutdown.
	<-ctx.Done()
	log.Info("Server: shutting down")
	err := server.Close()
	if err != nil {
		log.Errorf("Server: shutdown failed (%s)", err)
	}
}

// StartHttp starts the Web server in http mode.
func StartHttp(server *http.Server) {
	if err := server.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			log.Info("Server: shutdown complete")
		} else {
			log.Errorf("Server: %s", err)
		}
	}
}
