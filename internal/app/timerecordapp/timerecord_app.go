package timerecordapp

import (
	"context"

	"github.com/smart48ru/FaceIDApp/internal/domain"
)

type TimeRecordRepo interface {
	Save(ctx context.Context, u domain.TimeRecord) (uint64, error)
	GetLastByEmpoyeeID(ctx context.Context, id uint64) (domain.TimeRecord, error)
	List(ctx context.Context) ([]domain.TimeRecord, error)
}

type App struct {
	repo TimeRecordRepo
}

func New(repo TimeRecordRepo) *App {
	return &App{
		repo: repo,
	}
}
