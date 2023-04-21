package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/smart48ru/FaceIDApp/internal/config"
)

type Server struct {
	srv http.Server
}

func NewServer(conf *config.Config, h http.Handler) *Server {
	addr := fmt.Sprintf("%s:%d", conf.API.Host, conf.API.Port)

	s := &Server{}

	s.srv = http.Server{
		Addr:         addr,
		Handler:      h,
		ReadTimeout:  conf.API.ReadTimeout,
		WriteTimeout: conf.API.WriteTimeout,
	}
	return s
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	err := s.srv.Shutdown(ctx)
	if err != nil {
		log.Fatal().Msg("Server not shutdown")
	}
	cancel()
}

func (s *Server) Start() <-chan error {
	ret := make(chan error, 1)
	go func() {
		defer close(ret)

		err := s.srv.ListenAndServe()
		if err != http.ErrServerClosed {
			ret <- err
		}
	}()

	return ret
}
