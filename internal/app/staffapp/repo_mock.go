package staffapp

import (
	"context"

	"github.com/smart48ru/FaceIDApp/internal/domain"
	"github.com/stretchr/testify/mock"
)

type MockStaffRepo struct {
	mock.Mock
}

func (m *MockStaffRepo) Save(ctx context.Context, u domain.Employee) (uint64, error) {
	args := m.Called(ctx, u)
	return args.Get(0).(uint64), args.Error(1)
}

func (m *MockStaffRepo) Get(ctx context.Context, id uint64) (domain.Employee, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Employee), args.Error(1)
}

func (m *MockStaffRepo) Update(ctx context.Context, u domain.Employee) (domain.Employee, error) {
	args := m.Called(ctx, u)
	return args.Get(0).(domain.Employee), args.Error(1)
}

func (m *MockStaffRepo) Delete(ctx context.Context, id uint64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockStaffRepo) List(ctx context.Context) ([]domain.Employee, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.Employee), args.Error(1)
}
