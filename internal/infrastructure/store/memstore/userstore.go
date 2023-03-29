package memstore

import (
	"FaceIDApp/internal/entities/users"
	"context"
	"github.com/google/uuid"
	"sync"
)

type UserStore struct {
	sync.Mutex
	// TODO уточнить структуру хранения данных, заменить на актуальную
	m map[uuid.UUID]users.User
}

func (s *UserStore) Create(ctx context.Context, u users.User) (uuid.UUID, error) {
	//TODO implement me
	panic("implement me")
}

func (s *UserStore) Read(ctx context.Context, id uuid.UUID) (users.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *UserStore) Update(ctx context.Context, id uuid.UUID, u users.User) (users.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *UserStore) Delete(ctx context.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (s *UserStore) ReadAll(ctx context.Context) ([]users.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUseStore() *UserStore {
	return &UserStore{
		m: make(map[uuid.UUID]users.User),
	}
}
