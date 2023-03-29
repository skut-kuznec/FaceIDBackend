package service

import (
	"FaceIDApp/internal/usecases/app/repos/calendarrepo"
	"FaceIDApp/internal/usecases/app/repos/usersrepo"
)

type Store struct {
	u *usersrepo.Users
	c *calendarrepo.Calendar
}

func NewService(u *usersrepo.Users, c *calendarrepo.Calendar) *Store {
	r := &Store{
		u: u,
		c: c,
	}
	return r
}
