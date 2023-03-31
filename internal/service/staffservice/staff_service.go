package staffservice

import (
	"context"

	"github.com/smart48ru/FaceIDApp/internal/domain"
)

type StaffRepo interface {
	Create(ctx context.Context, u domain.Employee) (uint64, error)
	Read(ctx context.Context, id uint64) (domain.Employee, error)
	Update(ctx context.Context, id uint64, u domain.Employee) (domain.Employee, error)
	Delete(ctx context.Context, id uint64) error
	ReadAll(ctx context.Context) ([]domain.Employee, error)
}

type Service struct {
	repo StaffRepo
}

func New(repo StaffRepo) *Service {
	return &Service{
		repo: repo,
	}
}
