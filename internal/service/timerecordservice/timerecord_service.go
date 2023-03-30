package timerecordservice

import (
	"context"
	"time"

	"github.com/smart48ru/FaceIDApp/internal/domain"
)

type TimeRecordRepo interface {
	Create(ctx context.Context, c domain.TimeRecord) (int, error)
	Read(ctx context.Context, day time.Time) (domain.TimeRecord, error)
	Update(ctx context.Context, day time.Time, c domain.TimeRecord) (domain.Employee, error)
	Delete(ctx context.Context, day time.Time, userID int) error
	ReadAll(ctx context.Context) ([]domain.Employee, error)
}

type Service struct {
	repo TimeRecordRepo
}

func New(repo TimeRecordRepo) *Service {
	return &Service{
		repo: repo,
	}
}
