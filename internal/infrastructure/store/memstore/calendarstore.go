package memstore

import (
	"FaceIDApp/internal/entities/users"
	"FaceIDApp/internal/entities/workcalendar"
	"context"
	"sync"
	"time"

	"github.com/google/uuid"
)

type CalendarStore struct {
	sync.Mutex
	// TODO уточнить структуру хранения данных, заменить на актуальную
	m map[time.Time]workcalendar.WorkCalendar
}

func NewCalendarStore() *CalendarStore {
	return &CalendarStore{
		m: make(map[time.Time]workcalendar.WorkCalendar),
	}
}

func (c *CalendarStore) Create(ctx context.Context, cal workcalendar.WorkCalendar) (uuid.UUID, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CalendarStore) Read(ctx context.Context, day time.Time) (workcalendar.WorkCalendar, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CalendarStore) Update(ctx context.Context, day time.Time, cal workcalendar.WorkCalendar) (users.User, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CalendarStore) Delete(ctx context.Context, day time.Time, user uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (c *CalendarStore) ReadAll(ctx context.Context) ([]users.User, error) {
	//TODO implement me
	panic("implement me")
}
