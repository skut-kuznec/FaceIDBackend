package staffservice

import (
	"context"

	"github.com/smart48ru/FaceIDApp/internal/domain"
)

type StaffRepo interface {
	Create(ctx context.Context, u domain.Employee) (int, error)
	Read(ctx context.Context, id int) (domain.Employee, error)
	Update(ctx context.Context, id int, u domain.Employee) (domain.Employee, error)
	Delete(ctx context.Context, id int) error
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
