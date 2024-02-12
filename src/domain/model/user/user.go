package user

import (
	"homework/config"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/golang-jwt/jwt"
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
	UserName string `json:"userName"`
}

type UserListResponse struct {
	ID       string `json:"id"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	RoomID   string `json:"roomId"`
}

type Users []UserListResponse

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

func (u *User) GenerateToken(cnf config.Config) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": u.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(cnf.Seclet))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (u *Users) NewUserListResponse() []UserListResponse {
	var res []UserListResponse
	for _, user := range *u {
		res = append(res, UserListResponse{
			ID:       user.ID,
			UserName: user.UserName,
			Email:    user.Email,
			RoomID:   user.RoomID,
		})
	}
	return res
}
