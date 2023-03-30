package memstore

import (
	"context"
	"sync"
	"time"

	"github.com/smart48ru/FaceIDApp/internal/entities/stuff"
	"github.com/smart48ru/FaceIDApp/internal/entities/timerecord"

	"github.com/google/uuid"
)

type CalendarStore struct {
	sync.Mutex
	// TODO уточнить структуру хранения данных, заменить на актуальную
	m map[time.Time]timerecord.TimeRecord
}

func NewCalendarStore() *CalendarStore {
	return &CalendarStore{
		m: make(map[time.Time]timerecord.TimeRecord),
	}
}

func (c *CalendarStore) Create(ctx context.Context, cal timerecord.TimeRecord) (uuid.UUID, error) {
	panic("implement me")
}

func (c *CalendarStore) Read(ctx context.Context, day time.Time) (timerecord.TimeRecord, error) {
	panic("implement me")
}

func (c *CalendarStore) Update(ctx context.Context, day time.Time, cal timerecord.TimeRecord) (stuff.Stuff, error) {
	panic("implement me")
}

func (c *CalendarStore) Delete(ctx context.Context, day time.Time, user uuid.UUID) error {
	panic("implement me")
}

func (c *CalendarStore) ReadAll(ctx context.Context) ([]stuff.Stuff, error) {
	panic("implement me")
}
