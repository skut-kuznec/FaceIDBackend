package photorepo

import (
	"FaceIDApp/internal/entities/photo"
	"context"
	"github.com/google/uuid"
)

type PhotoStore interface {
	Create(ctx context.Context, ph photo.Photo) (uuid.UUID, error)
	Read(ctx context.Context, id uuid.UUID) (interface{}, error)
	ProcessedAI(ctx context.Context, ph photo.Photo) (int64, error)
	Delete(ctx context.Context, d uuid.UUID) error
}

type Photo struct {
	store PhotoStore
}

func NewPhotoRepo(store PhotoStore) *Photo {
	return &Photo{
		store: store,
	}
}
