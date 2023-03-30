package staffrepo

import (
	"context"
	"github.com/smart48ru/FaceIDApp/internal/domain"
	"reflect"
	"testing"
)

func TestRepo_Create(t *testing.T) {
	r := &Repo{
		m: make(map[int]domain.Employee),
	}
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

	// test creating an employee with an existing ID
	_, err = r.Create(ctx, u)
	if err == nil {
		t.Errorf("expected error creating employee with existing ID")
	}
	if len(r.m) != 1 {
		t.Errorf("expected map length 1, but got %d", len(r.m))
	}
}

func TestRepo_Delete(t *testing.T) {
	r := &Repo{
		m: make(map[int]domain.Employee),
	}
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
		t.Fatalf("unexpected error deleting employee: %v", err)
	}
}

func TestRepo_Read(t *testing.T) {
	r := &Repo{
		m: make(map[int]domain.Employee),
	}
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
