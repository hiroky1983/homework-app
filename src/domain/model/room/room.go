package room

import (
	"homework/websocket"
	"time"

	"github.com/uptrace/bun"
)

type Room struct {
	bun.BaseModel `bun:"table:room,alias:r"`

	ID        string    `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	CreatedAt time.Time `json:"created_at" bun:"default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" bun:"default:current_timestamp"`
	IsDeleted bool      `json:"is_deleted" bun:"default:false"`
}

type RoomMap struct {
	bun.BaseModel `bun:"table:room_map,alias:rm"`

	RoomID    string    `json:"room_id" bun:"room_id,type:uuid,notnull"`
	UserID    string    `json:"user_id" bun:"user_id,type:uuid,notnull"`
	CreatedAt time.Time `json:"created_at" bun:"default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" bun:"default:current_timestamp"`
}

var RoomToHub = map[uint]*websocket.Hub{}
