package staffapp

import (
	"context"

	"github.com/smart48ru/FaceIDApp/internal/domain"
)

type StaffRepo interface {
	Save(ctx context.Context, u domain.Employee) (uint64, error)
	Get(ctx context.Context, id uint64) (domain.Employee, error)
	Update(ctx context.Context, u domain.Employee) (domain.Employee, error)
	Delete(ctx context.Context, id uint64) error
	List(ctx context.Context) ([]domain.Employee, error)
}

type App struct {
	repo StaffRepo
}

func New(repo StaffRepo) *App {
	return &App{
		repo: repo,
	}
}
