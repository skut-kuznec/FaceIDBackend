package filestore

import (
	"FaceIDApp/internal/entities/photo"
	"context"

	"github.com/google/uuid"
)

type PhotoStore struct {
	dir string
}

func (p PhotoStore) Create(ctx context.Context, ph photo.Photo) (uuid.UUID, error) {
	//TODO implement me
	panic("implement me")
}

func (p PhotoStore) Read(ctx context.Context, id uuid.UUID) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (p PhotoStore) ProcessedAI(ctx context.Context, ph photo.Photo) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (p PhotoStore) Delete(ctx context.Context, d uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func NewPhotoStore(dir string) *PhotoStore {
	return &PhotoStore{
		dir: dir,
	}
}
