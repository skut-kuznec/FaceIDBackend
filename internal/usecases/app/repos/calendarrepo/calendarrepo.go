package calendarrepo

import (
	"FaceIDApp/internal/entities/users"
	"FaceIDApp/internal/entities/workcalendar"
	"context"
	"time"

	"github.com/google/uuid"
)

type CalendarStorage interface {
	Create(ctx context.Context, c workcalendar.WorkCalendar) (uuid.UUID, error)
	Read(ctx context.Context, day time.Time) (workcalendar.WorkCalendar, error)
	Update(ctx context.Context, day time.Time, c workcalendar.WorkCalendar) (users.User, error)
	Delete(ctx context.Context, day time.Time, user uuid.UUID) error
	ReadAll(ctx context.Context) ([]users.User, error)
}

type Calendar struct {
	store CalendarStorage
}

func NewCalendarRepo(store CalendarStorage) *Calendar {
	return &Calendar{
		store: store,
	}
}
