package filestore

import (
	"context"

	"github.com/smart48ru/FaceIDApp/internal/entities/image"

	"github.com/google/uuid"
)

type PhotoStore struct {
	dir string
}

func (p PhotoStore) Create(ctx context.Context, ph image.Image) (uuid.UUID, error) {
	panic("implement me")
}

func (p PhotoStore) Read(ctx context.Context, id uuid.UUID) (interface{}, error) {
	panic("implement me")
}

func (p PhotoStore) ProcessedAI(ctx context.Context, ph image.Image) (int64, error) {
	panic("implement me")
}

func (p PhotoStore) Delete(ctx context.Context, d uuid.UUID) error {
	panic("implement me")
}

func NewPhotoStore(dir string) *PhotoStore {
	return &PhotoStore{
		dir: dir,
	}
}
