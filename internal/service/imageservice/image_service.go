package imageservice

import (
	"context"

	"github.com/google/uuid"
	"github.com/smart48ru/FaceIDApp/internal/domain"
)

type ImageRepo interface {
	Create(ctx context.Context, ph domain.Image) (int, error)
	Read(ctx context.Context, id int) (domain.Image, error)
	Delete(ctx context.Context, d uuid.UUID) error
}

type Service struct {
	repo ImageRepo
}

func New(repo ImageRepo) *Service {
	return &Service{
		repo: repo,
	}
}
