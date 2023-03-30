package service

import (
	"github.com/smart48ru/FaceIDApp/internal/usecases/app/repos/imagerepo"
	"github.com/smart48ru/FaceIDApp/internal/usecases/app/repos/stuffrepo"
	"github.com/smart48ru/FaceIDApp/internal/usecases/app/repos/timerecordrepo"
)

type Store struct {
	u *stuffrepo.Stuff
	c *timerecordrepo.TimeRecord
	p *imagerepo.Image
}

func NewService(u *stuffrepo.Stuff, c *timerecordrepo.TimeRecord, p *imagerepo.Image) *Store {
	r := &Store{
		u: u,
		c: c,
		p: p,
	}
	return r
}
