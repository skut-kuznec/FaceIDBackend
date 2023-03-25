package memstore

import (
	"context"
	"sync"
)

type Store struct {
	sync.Mutex
	// TODO уточнить структуру хранения данных, заменить на актуальную
	m map[interface{}]interface{}
}

func (s *Store) Create(ctx context.Context, u interface{}) (interface{}, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) Read(ctx context.Context, u interface{}) (interface{}, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) Delete(ctx context.Context, u interface{}) error {
	// TODO implement me
	panic("implement me")
}

func (s *Store) SearchUsers(ctx context.Context, u interface{}) (interface{}, error) {
	// TODO implement me
	panic("implement me")
}

func NewMemStore() *Store {
	return &Store{
		m: make(map[interface{}]interface{}),
	}
}
