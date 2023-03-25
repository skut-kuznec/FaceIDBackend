package usersrepo

import (
	"context"
)

type UserStore interface {
	Create(ctx context.Context, u interface{}) (interface{}, error)
	Read(ctx context.Context, u interface{}) (interface{}, error)
	Delete(ctx context.Context, u interface{}) error
	SearchUsers(ctx context.Context, u interface{}) (interface{}, error)
}

type Users struct {
	store UserStore
}

func NewUsersRepo(store UserStore) *Users {
	return &Users{
		store: store,
	}
}
