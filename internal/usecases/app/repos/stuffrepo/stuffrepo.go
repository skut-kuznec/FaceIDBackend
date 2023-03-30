package stuffrepo

import (
	"context"

	"FaceIDApp/internal/entities/stuff"

	"github.com/google/uuid"
)

type StuffStore interface {
	Create(ctx context.Context, u stuff.Stuff) (uuid.UUID, error)
	Read(ctx context.Context, id uuid.UUID) (stuff.Stuff, error)
	Update(ctx context.Context, id uuid.UUID, u stuff.Stuff) (stuff.Stuff, error)
	Delete(ctx context.Context, id uuid.UUID) error
	ReadAll(ctx context.Context) ([]stuff.Stuff, error)
}

type Stuff struct {
	store StuffStore
}

func NewUsersRepo(store StuffStore) *Stuff {
	return &Stuff{
		store: store,
	}
}
