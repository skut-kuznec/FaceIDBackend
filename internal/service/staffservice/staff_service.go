package staffservice

import (
	"context"

	"github.com/smart48ru/FaceIDApp/internal/domain"
)

type StaffRepo interface {
	Create(ctx context.Context, u domain.Staff) (int, error)
	Read(ctx context.Context, id int) (domain.Staff, error)
	Update(ctx context.Context, id int, u domain.Staff) (domain.Staff, error)
	Delete(ctx context.Context, id int) error
	ReadAll(ctx context.Context) ([]domain.Staff, error)
}

type Service struct {
	repo StaffRepo
}

func New(repo StaffRepo) *Service {
	return &Service{
		repo: repo,
	}
}
