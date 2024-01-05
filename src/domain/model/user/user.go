package user

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:user,alias:u"`

	ID uint `bun:"id,pk,type:uuid"`
	Email string `bun:"type:varchar(255)"`
	Password string `bun:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID uint `json:"id" bun:"primary_key"` 
	Email string `json:"email" bun:"unique"`
}
