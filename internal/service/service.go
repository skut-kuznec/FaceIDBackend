package service

import (
	"FaceIDApp/internal/usecases/app/repos/calendarrepo"
	"FaceIDApp/internal/usecases/app/repos/photorepo"
	"FaceIDApp/internal/usecases/app/repos/usersrepo"
)

type Store struct {
	u *usersrepo.Users
	c *calendarrepo.Calendar
	p *photorepo.Photo
}

func NewService(u *usersrepo.Users, c *calendarrepo.Calendar, p *photorepo.Photo) *Store {
	r := &Store{
		u: u,
		c: c,
		p: p,
	}
	return r
}
