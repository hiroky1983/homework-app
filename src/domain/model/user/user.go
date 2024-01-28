package user

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:user,alias:u"`

	ID         string    `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	UserName   string    `json:"user_name" bun:"type:varchar(255)"`
	Email      string    `json:"email" bun:"type:varchar(255)"`
	Password   string    `json:"password" bun:"type:varchar(255)"`
	ImagePath  string    `json:"image_path" bun:"type:varchar(255)"`
	IsVerified bool      `json:"Is_verified" bun:"default:false"`
	GoogleID   string    `json:"google_id" bun:"type:varchar(255)"`
	CreatedAt  time.Time `json:"created_at" bun:"default:current_timestamp"`
	UpdatedAt  time.Time `json:"updated_at" bun:"default:current_timestamp"`
	IsDeleted  bool      `json:"is_deleted" bun:"default:false"`
}

type UserResponse struct {
	ID    string `json:"id" bun:"primary_key"`
	Email string `json:"email" bun:"unique"`
}

type LonginResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type UserProfileRequest struct {
	UserName  string `json:"userName"`
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
