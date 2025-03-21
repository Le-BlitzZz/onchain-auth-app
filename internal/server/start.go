package server

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Start(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
		}
	}()

	router := gin.New()

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go StartHttp(server)

	<-ctx.Done()
	log.Info("Server: shutting down")
	err := server.Close()
	if err != nil {
		log.Errorf("Server: shutdown failed (%s)", err)
	}
}

func StartHttp(server *http.Server) {
	if err := server.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			log.Info("Server: shutdown complete")
		} else {
			log.Errorf("Server: %s", err)
		}
	}
}
