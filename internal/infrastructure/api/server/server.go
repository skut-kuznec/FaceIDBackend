package server

import (
	"FaceIDApp/internal/config"
	"FaceIDApp/internal/service"
	"context"
	"fmt"
	"net/http"
	"time"
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
	s.srv.Shutdown(ctx)
	cancel()
}
func (s *Server) Start(eap *service.Store) {
	s.service = eap
	go s.srv.ListenAndServe()
}
