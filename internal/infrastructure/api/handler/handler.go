package handler

import "FaceIDApp/internal/usecases/repos/usersrepo"

type Handlers struct {
	repo *usersrepo.Users
}

func NewHandler(r *usersrepo.Users) *Handlers {
	h := &Handlers{
		repo: r,
	}
	return h
}
