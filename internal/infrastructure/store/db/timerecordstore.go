package db

import (
	"context"
	"time"

	"github.com/smart48ru/FaceIDApp/internal/entities/stuff"
	"github.com/smart48ru/FaceIDApp/internal/entities/timerecord"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type TimeRecordStore struct {
	db *sqlx.DB
}

func NewTimeRecordStore(db *sqlx.DB) *TimeRecordStore {
	us := &TimeRecordStore{
		db: db,
	}
	return us
}

func (c *TimeRecordStore) Create(ctx context.Context, cal timerecord.TimeRecord) (uuid.UUID, error) {
	panic("implement me")
}

func (c *TimeRecordStore) Read(ctx context.Context, day time.Time) (timerecord.TimeRecord, error) {
	panic("implement me")
}

func (c *TimeRecordStore) Update(ctx context.Context, day time.Time, cal timerecord.TimeRecord) (stuff.Stuff, error) {
	panic("implement me")
}

func (c *TimeRecordStore) Delete(ctx context.Context, day time.Time, user uuid.UUID) error {
	panic("implement me")
}

func (c *TimeRecordStore) ReadAll(ctx context.Context) ([]stuff.Stuff, error) {
	panic("implement me")
}
