package user

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:user,alias:u"`

	ID        string    `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	Email     string    `bun:"type:varchar(255),unique"`
	Password  string    `bun:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID    string `json:"id" bun:"primary_key"`
	Email string `json:"email" bun:"unique"`
}

type LonginResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(
			&u.Email,
			validation.Required.Error("email is required"),
			validation.RuneLength(1, 30).Error("limited max 30 char"),
			is.Email.Error("is not valid email format"),
		),
		validation.Field(
			&u.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(8, 30).Error("limited min 8 max 30 char"),
		),
	)
}
