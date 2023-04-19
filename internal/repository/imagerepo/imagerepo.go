package imagerepo

import (
	"context"
	"io"

	"github.com/smart48ru/FaceIDApp/internal/app/imageapp"
	"github.com/smart48ru/FaceIDApp/internal/domain"
)

var _ imageapp.ImageRepo = &Repo{}

type Repo struct{}

func New() *Repo {
	return &Repo{}
}

// Get implements imageapp.ImageRepo
func (re *Repo) Get(ctx context.Context, id uint64) (io.Reader, error) {
	panic("unimplemented")
}

// GetDescByID implements imageapp.ImageRepo
func (re *Repo) GetDescByID(ctx context.Context, in uint64) (domain.Image, error) {
	panic("unimplemented")
}

// Save implements imageapp.ImageRepo
func (re *Repo) Save(ctx context.Context, img []byte, descr domain.FaceDescriptor) (domain.Image, error) {
	panic("unimplemented")
}

func (re *Repo) List(ctx context.Context) ([]domain.Image, error) {
	panic("unimplemented")
}
