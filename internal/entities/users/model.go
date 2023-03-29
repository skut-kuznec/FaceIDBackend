package users

import (
	"FaceIDApp/internal/usecases/app/repos/photorepo"

	"github.com/google/uuid"
)

type User struct {
	ID      uuid.UUID       `json:"id,omitempty"`
	AID     int64           `json:"aid,omitempty"`
	Name    string          `json:"name,omitempty"`
	Surname string          `json:"surname,omitempty"`
	Photo   photorepo.Photo `json:"photo,omitempty"`
	Block   bool            `json:"block,omitempty"`
	Admin   bool            `json:"admin,omitempty"`
}
