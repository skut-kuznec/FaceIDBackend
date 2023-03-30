package imagerepo

import (
	"context"

	"github.com/smart48ru/FaceIDApp/internal/entities/image"

	"github.com/google/uuid"
)

type ImageStore interface {
	Create(ctx context.Context, ph image.Image) (uuid.UUID, error)
	Read(ctx context.Context, id uuid.UUID) (interface{}, error)
	ProcessedAI(ctx context.Context, ph image.Image) (int64, error)
	Delete(ctx context.Context, d uuid.UUID) error
}

type Image struct {
	store ImageStore
}

func NewPhotoRepo(store ImageStore) *Image {
	return &Image{
		store: store,
	}
}
