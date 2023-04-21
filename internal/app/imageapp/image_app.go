package imageapp

import (
	"context"
	"io"

	"github.com/smart48ru/FaceIDApp/internal/domain"
)

type ImageRepo interface {
	Save(ctx context.Context, img []byte, descr domain.FaceDescriptor) (domain.Image, error)
	Get(ctx context.Context, id uint64) (io.Reader, error)
	GetDescByID(ctx context.Context, in uint64) (domain.Image, error)
}

type FaceService interface {
	GetFaceDescriptor(img []byte) (domain.FaceDescriptor, error)
	SearchImageID(descr domain.FaceDescriptor, images []domain.Image) (uint64, error)
}

type App struct {
	repo        ImageRepo
	faceService FaceService
}

func New(repo ImageRepo, faceService FaceService) *App {
	return &App{
		repo:        repo,
		faceService: faceService,
	}
}

func (a *App) Save(ctx context.Context, r io.Reader) (domain.Image, error) {
	img, err := io.ReadAll(r)
	if err != nil {
		return domain.Image{}, err
	}

	descr, err := a.faceService.GetFaceDescriptor(img)
	if err != nil {
		return domain.Image{}, err
	}

	ret, err := a.repo.Save(ctx, img, descr)
	if err != nil {
		return domain.Image{}, err
	}

	return ret, nil
}
