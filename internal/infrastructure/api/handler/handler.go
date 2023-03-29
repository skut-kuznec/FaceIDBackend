package handler

import (
	"FaceIDApp/internal/service"
)

type Handlers struct {
	s *service.Store
}

func NewHandler(s *service.Store) *Handlers {
	h := &Handlers{
		s: s,
	}
	return h
}
