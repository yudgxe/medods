package server

import (
	"context"
	"fmt"
	"medods/logger"
	"medods/utils"
	"net/http"
	"os"
	"os/signal"
)

type Server struct {
	*http.Server
}

func New(host, port string, router http.Handler) *Server {
	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: router,
	}

	return &Server{&server}
}

func (s *Server) Start() {
	go func() {
		logger.Log().Debugf("server start at: %s", s.Addr)

		if err := s.ListenAndServe(); err != http.ErrServerClosed && err != nil {
			utils.Panicf("error on server listen - %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	if err := s.Shutdown(context.Background()); err != nil {
		logger.Log().Errorf("error on server shutdown - %s", err)
	}

	logger.Log().Debug("server gracefully shutdown")
}
