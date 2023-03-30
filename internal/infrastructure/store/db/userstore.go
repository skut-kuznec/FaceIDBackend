package db

import (
	"context"

	"FaceIDApp/internal/entities/stuff"

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

func (u *UserStore) Create(ctx context.Context, us stuff.Stuff) (uuid.UUID, error) {
	panic("implement me")
}

func (u *UserStore) Read(ctx context.Context, id uuid.UUID) (stuff.Stuff, error) {
	panic("implement me")
}

func (u *UserStore) Update(ctx context.Context, id uuid.UUID, us stuff.Stuff) (stuff.Stuff, error) {
	panic("implement me")
}

func (u *UserStore) Delete(ctx context.Context, id uuid.UUID) error {
	panic("implement me")
}

func (u *UserStore) ReadAll(ctx context.Context) ([]stuff.Stuff, error) {
	panic("implement me")
}

func (u *UserStore) ReadIAID(ctx context.Context, aiID int64) (stuff.Stuff, error) {
	panic("implement me")
}
