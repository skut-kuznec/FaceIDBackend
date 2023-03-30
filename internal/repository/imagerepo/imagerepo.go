package imagerepo

import (
	"context"

	"github.com/google/uuid"
	"github.com/smart48ru/FaceIDApp/internal/domain"
	"github.com/smart48ru/FaceIDApp/internal/service/imageservice"
)

var _ imageservice.ImageRepo = &Repo{}

type Repo struct {
}

// Create implements imageservice.ImageRepo
func (*Repo) Create(ctx context.Context, ph domain.Image) (int, error) {
	panic("unimplemented")
}

// Delete implements imageservice.ImageRepo
func (*Repo) Delete(ctx context.Context, d uuid.UUID) error {
	panic("unimplemented")
}

// Read implements imageservice.ImageRepo
func (*Repo) Read(ctx context.Context, id int) (domain.Image, error) {
	panic("unimplemented")
}

func New() *Repo {
	return &Repo{}
}
