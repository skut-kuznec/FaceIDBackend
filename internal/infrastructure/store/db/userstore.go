package db

import (
	"FaceIDApp/internal/entities/users"
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserStore struct {
	db *sqlx.DB
}

func NewUSerStore(db *sqlx.DB) *UserStore {
	us := &UserStore{
		db: db,
	}
	return us
}

func (u *UserStore) Create(ctx context.Context, us users.User) (uuid.UUID, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserStore) Read(ctx context.Context, id uuid.UUID) (users.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserStore) Update(ctx context.Context, id uuid.UUID, us users.User) (users.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserStore) Delete(ctx context.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserStore) ReadAll(ctx context.Context) ([]users.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserStore) ReadIAID(ctx context.Context, aiID int64) (users.User, error) {
	//TODO implement me
	panic("implement me")
}
