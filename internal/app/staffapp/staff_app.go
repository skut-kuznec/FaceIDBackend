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

func (a *App) AddEmployee(ctx context.Context, name string, photoID int) (uint64, error) {
	employee := domain.Employee{
		Name:    name,
		PhotoID: photoID,
	}
	id, err := a.repo.Save(ctx, employee)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (a *App) GetEmployee(ctx context.Context, id uint64) (domain.Employee, error) {
	employee, err := a.repo.Get(ctx, id)
	if err != nil {
		return domain.Employee{}, err
	}
	return employee, nil
}

func (a *App) UpdateEmployee(ctx context.Context, employee domain.Employee) (domain.Employee, error) {
	updateEmployee, err := a.repo.Update(ctx, employee)
	if err != nil {
		return domain.Employee{}, err
	}
	return updateEmployee, err
}

func (a *App) DeleteEmployee(ctx context.Context, id uint64) error {
	err := a.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) ListEmployee(ctx context.Context) ([]domain.Employee, error) {
	list, err := a.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return list, nil
}
