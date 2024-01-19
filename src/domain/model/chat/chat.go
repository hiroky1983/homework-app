package chat

import (
	"time"

	"github.com/uptrace/bun"
)

type Chat struct {
	bun.BaseModel `bun:"table:chat,alias:c"`

	ID        string    `json:"id" bun:"id,pk,type:int,autoincrement"`
	Message 	string    `json:"message" bun:"type:varchar(255),notnull"`
	UserID    string    `json:"user_id" bun:"type:uuid,notnull"`
	CreatedAt time.Time `json:"created_at" bun:"default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" bun:"default:current_timestamp"`
	IsDeleted bool      `json:"is_deleted" bun:"default:false"`
}
