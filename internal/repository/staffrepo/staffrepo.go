package staffrepo

import (
	"context"

	"github.com/smart48ru/FaceIDApp/internal/domain"
	"github.com/smart48ru/FaceIDApp/internal/service/staffservice"
)

var _ staffservice.StaffRepo = &Repo{}

type Repo struct {
}

// Create implements staffservice.StaffRepo
func (*Repo) Create(ctx context.Context, u domain.Staff) (int, error) {
	panic("unimplemented")
}

// Delete implements staffservice.StaffRepo
func (*Repo) Delete(ctx context.Context, id int) error {
	panic("unimplemented")
}

// Read implements staffservice.StaffRepo
func (*Repo) Read(ctx context.Context, id int) (domain.Staff, error) {
	panic("unimplemented")
}

// ReadAll implements staffservice.StaffRepo
func (*Repo) ReadAll(ctx context.Context) ([]domain.Staff, error) {
	panic("unimplemented")
}

// Update implements staffservice.StaffRepo
func (*Repo) Update(ctx context.Context, id int, u domain.Staff) (domain.Staff, error) {
	panic("unimplemented")
}

func New() *Repo {
	return &Repo{}
}
