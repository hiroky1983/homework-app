package room

import (
	"time"

	"github.com/uptrace/bun"
)

type Room struct {
	bun.BaseModel `bun:"table:room,alias:r"`

	RoomID    uint64    `json:"id" bun:"id,pk,autoincrement"`
	UserID    string    `json:"user_id" bun:"type:uuid,notnull"`
	CreatedAt time.Time `json:"created_at" bun:"default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" bun:"default:current_timestamp"`
	IsDeleted bool      `json:"is_deleted" bun:"default:false"`
}