package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/smart48ru/FaceIDApp/internal/config"
	"github.com/smart48ru/FaceIDApp/internal/service"

	"github.com/rs/zerolog/log"
)

type Server struct {
	srv     http.Server
	service *service.Store
}

func NewServer(conf config.API, h http.Handler) *Server {
	addr := fmt.Sprintf("%s:%d", conf.APIHost(), conf.APIPort())

	s := &Server{}

	s.srv = http.Server{
		Addr:              addr,
		Handler:           h,
		ReadTimeout:       time.Duration(conf.APIReadTimeOut()) * time.Second,
		WriteTimeout:      time.Duration(conf.APIWriteTimeOut()) * time.Second,
		ReadHeaderTimeout: time.Duration(conf.APIReadHeadTimeOut()) * time.Second,
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

func (s *Server) Start() {

	go s.srv.ListenAndServe()
}
