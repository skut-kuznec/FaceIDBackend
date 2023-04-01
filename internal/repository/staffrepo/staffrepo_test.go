package staffrepo

import (
	"context"
	"reflect"
	"testing"

	"github.com/smart48ru/FaceIDApp/internal/domain"
)

func TestRepo_Create(t *testing.T) {
	r := New()

	u := domain.Employee{
		ID:      1,
		Name:    "John Doe",
		PhotoID: 0,
	}
	ctx := context.Background()

	// test creating a new employee
	id, err := r.Create(ctx, u)
	if err != nil {
		t.Fatalf("unexpected error creating employee: %v", err)
	}
	if id != u.ID {
		t.Errorf("expected employee ID %d, but got %d", u.ID, id)
	}
	if len(r.m) != 1 {
		t.Errorf("expected map length 1, but got %d", len(r.m))
	}
	if e, ok := r.m[u.ID]; !ok {
		t.Errorf("expected employee with ID %d to exist in map", u.ID)
	} else if !reflect.DeepEqual(e, u) {
		t.Errorf("expected employee %v, but got %v", u, e)
	}
}

func TestRepo_Delete(t *testing.T) {
	r := New()
	u := domain.Employee{
		ID:      123,
		Name:    "John Doe",
		PhotoID: 0,
	}
	ctx := context.Background()

	// test creating a new employee
	id, err := r.Create(ctx, u)
	if err != nil {
		t.Fatalf("unexpected error creating employee: %v", err)
	}
	// deleting employee
	err = r.Delete(ctx, id)
	if err != nil {
		t.Fatalf("unexpected error deleting employee: %v", err)
	}
	// delete again employee
	err = r.Delete(ctx, id)
	if err == nil {
		t.Fatalf("expected error when deleting already deleted employee, but got nil: %v", err)
	}
}

func TestRepo_Read(t *testing.T) {
	r := New()
	u := domain.Employee{
		ID:      123,
		Name:    "John Doe",
		PhotoID: 0,
	}
	ctx := context.Background()

	// test creating a new employee
	id, err := r.Create(ctx, u)
	if err != nil {
		t.Fatalf("unexpected error creating employee: %v", err)
	}
	// read employee
	_, err = r.Read(ctx, id)
	if err != nil {
		t.Fatalf("unexpected error read employee: %v", err)
	}
	// read not create employee
	_, err = r.Read(ctx, id*2)
	if err == nil {
		t.Fatalf("unexpected error read employee: %v", err)
	}
}

func TestRepo_Update(t *testing.T) {
	r := New()
	u := domain.Employee{
		ID:      123,
		Name:    "John Doe",
		PhotoID: 0,
	}
	ctx := context.Background()

	// test creating a new employee
	id, err := r.Create(ctx, u)
	if err != nil {
		t.Fatalf("unexpected error creating employee: %v", err)
	}
	u.Name = "NEW John Doe"
	u.PhotoID = 10
	emp, err := r.Update(ctx, id, u)
	if err != nil {
		t.Fatalf("unexpected error update employee: %v", err)
	}
	if u.Name != emp.Name {
		t.Errorf("expected employee name %s, but got %s", u.Name, emp.Name)
	}

	if !reflect.DeepEqual(emp, u) {
		t.Errorf("expected employee %v, but got %v", u, emp)
	}
}

func TestRepo_ReadAll(t *testing.T) {
	r := New()

	employees := []domain.Employee{
		{
			ID:      1,
			Name:    "John Doe",
			PhotoID: 0,
		},
		{
			ID:      2,
			Name:    "Verlom terkov",
			PhotoID: 1,
		},
		{
			ID:      3,
			Name:    "Piter Parker",
			PhotoID: 12,
		},
	}

	ctx := context.Background()

	// test creating a new employee
	for _, emp := range employees {
		if _, err := r.Create(ctx, emp); err != nil {
			t.Fatalf("unexpected error creating employee: %s", err)
		}
	}

	savedEmpoloyees, err := r.ReadAll(ctx)
	if err != nil {
		t.Fatalf("unexpected error read all employee: %v", err)
	}
	if !reflect.DeepEqual(employees, savedEmpoloyees) {
		t.Errorf("expected employee %v, but got %v", employees, savedEmpoloyees)
	}
	if len(employees) != len(savedEmpoloyees) {
		t.Errorf("expected len employee %v, but got len %v", len(employees), len(savedEmpoloyees))
	}
}
