package db

import (
	"FaceIDApp/internal/entities/users"
	"FaceIDApp/internal/entities/workcalendar"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type CalendarStore struct {
	db *sqlx.DB
}

func NewCalendarStore(db *sqlx.DB) *CalendarStore {
	us := &CalendarStore{
		db: db,
	}
	return us
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
