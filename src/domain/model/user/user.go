package user

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:user,alias:u"`

	ID uint `json:"id" bun:"primary_key"`
	Email string `json:"email" bun:"unique"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID uint `json:"id" bun:"primary_key"` 
	Email string `json:"email" bun:"unique"`
}
