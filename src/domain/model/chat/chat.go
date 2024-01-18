package chat

import (
	"time"

	"github.com/uptrace/bun"
)

type Chat struct {
	bun.BaseModel `bun:"table:chat,alias:c"`

	ID        string    `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	UserName 	string    `bun:"type:varchar(255)"`
	Email     string    `bun:"type:varchar(255)"`
	Password  string    `bun:"type:varchar(255)"`
	ImagePath string    `bun:"type:varchar(255)"`
	IsVerified bool     `bun:"default:false"`
	GoogleID  string    `bun:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at" bun:"default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" bun:"default:current_timestamp"`
	IsDeleted bool      `json:"is_deleted" bun:"default:false"`
}