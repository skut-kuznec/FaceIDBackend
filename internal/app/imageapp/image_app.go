package imageapp

import (
	"context"
	"io"

	"github.com/smart48ru/FaceIDApp/internal/domain"
)

type ImageRepo interface {
	Save(ctx context.Context, r io.Reader) (domain.Image, error)
	Get(ctx context.Context, id uint64) (io.Reader, error)
	GetDescByID(ctx context.Context, in uint64) (domain.Image, error)
}

type App struct {
	repo ImageRepo
}

func New(repo ImageRepo) *App {
	return &App{
		repo: repo,
	}
}
