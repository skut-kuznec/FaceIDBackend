package memstore

import (
	"context"
	"sync"

	"github.com/smart48ru/FaceIDApp/internal/entities/stuff"

	"github.com/google/uuid"
)

type UserStore struct {
	sync.Mutex
	// TODO уточнить структуру хранения данных, заменить на актуальную
	m map[uuid.UUID]stuff.Stuff
}

func (s *UserStore) Create(ctx context.Context, u stuff.Stuff) (uuid.UUID, error) {
	panic("implement me")
}

func (s *UserStore) Read(ctx context.Context, id uuid.UUID) (stuff.Stuff, error) {
	panic("implement me")
}

func (s *UserStore) Update(ctx context.Context, id uuid.UUID, u stuff.Stuff) (stuff.Stuff, error) {
	panic("implement me")
}

func (s *UserStore) Delete(ctx context.Context, id uuid.UUID) error {
	panic("implement me")
}

func (s *UserStore) ReadAll(ctx context.Context) ([]stuff.Stuff, error) {
	panic("implement me")
}

func NewUseStore() *UserStore {
	return &UserStore{
		m: make(map[uuid.UUID]stuff.Stuff),
	}
}
