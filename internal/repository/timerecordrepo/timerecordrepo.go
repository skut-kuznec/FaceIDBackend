package timerecordrepo

import (
	"context"

	"github.com/smart48ru/FaceIDApp/internal/app/timerecordapp"
	"github.com/smart48ru/FaceIDApp/internal/domain"
)

var _ timerecordapp.TimeRecordRepo = &Repo{}

type Repo struct{}

func New() *Repo {
	return &Repo{}
}

// GetLastByEmpoyeeID implements timerecordapp.TimeRecordRepo
func (r *Repo) GetLastByEmpoyeeID(ctx context.Context, id uint64) (domain.TimeRecord, error) {
	panic("unimplemented")
}

// List implements timerecordapp.TimeRecordRepo
func (r *Repo) List(ctx context.Context) ([]domain.TimeRecord, error) {
	panic("unimplemented")
}

// Save implements timerecordapp.TimeRecordRepo
func (r *Repo) Save(ctx context.Context, u domain.TimeRecord) (uint64, error) {
	panic("unimplemented")
}
