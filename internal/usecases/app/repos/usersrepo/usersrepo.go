package usersrepo

import (
	"FaceIDApp/internal/entities/users"
	"context"

	"github.com/google/uuid"
)

type UserStore interface {
	Create(ctx context.Context, u users.User) (uuid.UUID, error)
	Read(ctx context.Context, id uuid.UUID) (users.User, error)
	Update(ctx context.Context, id uuid.UUID, u users.User) (users.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
	ReadAll(ctx context.Context) ([]users.User, error)
}

type Users struct {
	store UserStore
}

func NewUsersRepo(store UserStore) *Users {
	return &Users{
		store: store,
	}
}
