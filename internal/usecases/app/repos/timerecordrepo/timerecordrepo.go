package timerecordrepo

import (
	"context"
	"time"

	"github.com/smart48ru/FaceIDApp/internal/entities/stuff"
	"github.com/smart48ru/FaceIDApp/internal/entities/timerecord"

	"github.com/google/uuid"
)

type TimeRecordStorage interface {
	Create(ctx context.Context, c timerecord.TimeRecord) (uuid.UUID, error)
	Read(ctx context.Context, day time.Time) (timerecord.TimeRecord, error)
	Update(ctx context.Context, day time.Time, c timerecord.TimeRecord) (stuff.Stuff, error)
	Delete(ctx context.Context, day time.Time, user uuid.UUID) error
	ReadAll(ctx context.Context) ([]stuff.Stuff, error)
}

type TimeRecord struct {
	store TimeRecordStorage
}

func NewCalendarRepo(store TimeRecordStorage) *TimeRecord {
	return &TimeRecord{
		store: store,
	}
}
