package timerecordrepo

import (
	"context"
	"time"

	"github.com/smart48ru/FaceIDApp/internal/domain"
	"github.com/smart48ru/FaceIDApp/internal/service/timerecordservice"
)

var _ timerecordservice.TimeRecordRepo = &Repo{}

type Repo struct{}

// Create implements timerecordservice.TimeRecordRepo
func (*Repo) Create(ctx context.Context, c domain.TimeRecord) (int, error) {
	panic("unimplemented")
}

// Delete implements timerecordservice.TimeRecordRepo
func (*Repo) Delete(ctx context.Context, day time.Time, userID int) error {
	panic("unimplemented")
}

// Read implements timerecordservice.TimeRecordRepo
func (*Repo) Read(ctx context.Context, day time.Time) (domain.TimeRecord, error) {
	panic("unimplemented")
}

// ReadAll implements timerecordservice.TimeRecordRepo
func (*Repo) ReadAll(ctx context.Context) ([]domain.Employee, error) {
	panic("unimplemented")
}

// Update implements timerecordservice.TimeRecordRepo
func (*Repo) Update(ctx context.Context, day time.Time, c domain.TimeRecord) (domain.Employee, error) {
	panic("unimplemented")
}

func New() *Repo {
	return &Repo{}
}
