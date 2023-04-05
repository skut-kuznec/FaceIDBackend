package staffrepo

import (
	"context"
	"fmt"
	"sort"
	"sync"

	"github.com/smart48ru/FaceIDApp/internal/app/staffapp"
	"github.com/smart48ru/FaceIDApp/internal/domain"
)

var _ staffapp.StaffRepo = &Repo{}

type Repo struct {
	mu  sync.Mutex
	m   map[uint64]domain.Employee
	seq uint64
}

func New() *Repo {
	return &Repo{
		m: make(map[uint64]domain.Employee),
	}
}

// Save implements staffservice.StaffRepo
func (r *Repo) Save(ctx context.Context, e domain.Employee) (uint64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
	}
	r.seq += 1
	e.ID = r.seq
	r.m[r.seq] = e

	return r.seq, nil
}

// Delete implements staffservice.StaffRepo
func (r *Repo) Delete(ctx context.Context, id uint64) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	_, ok := r.m[id]
	if !ok {
		return fmt.Errorf("employee id=%d not exist", id)
	}
	delete(r.m, id)

	return nil
}

// Get implements staffservice.StaffRepo
func (r *Repo) Get(ctx context.Context, id uint64) (domain.Employee, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	select {
	case <-ctx.Done():
		return domain.Employee{}, ctx.Err()
	default:
	}

	if v, ok := r.m[id]; ok {
		return v, nil
	}

	return domain.Employee{}, fmt.Errorf("employee id=%d not found", id)
}

// List implements staffservice.StaffRepo
func (r *Repo) List(ctx context.Context) ([]domain.Employee, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	employees := make([]domain.Employee, 0, len(r.m))
	for _, employee := range r.m {
		employees = append(employees, employee)
	}
	sort.Slice(employees, func(i, j int) bool {
		return employees[i].ID < employees[j].ID
	})

	return employees, nil
}

// Update implements staffservice.StaffRepo
func (r *Repo) Update(ctx context.Context, u domain.Employee) (domain.Employee, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	select {
	case <-ctx.Done():
		return domain.Employee{}, ctx.Err()
	default:
	}

	if _, ok := r.m[u.ID]; !ok {
		return domain.Employee{}, fmt.Errorf("employee with id = %d not found", u.ID)
	}
	r.m[u.ID] = u

	return u, nil
}
