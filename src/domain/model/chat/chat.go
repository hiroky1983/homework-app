package chat

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/uptrace/bun"
)

const (
	SenderMe    = "me"
	SenderOther = "other"
)

type Chat struct {
	bun.BaseModel `bun:"table:chat,alias:c"`

	ID        uint64    `json:"id" bun:"id,pk,autoincrement"`
	Message   string    `json:"message" bun:"type:varchar(255),notnull"`
	UserID    string    `json:"user_id" bun:"type:uuid,notnull"`
	RoomID    uint64    `json:"room_id" bun:"type:uuid,notnull"`
	CreatedAt time.Time `json:"created_at" bun:"default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" bun:"default:current_timestamp"`
	IsDeleted bool      `json:"is_deleted" bun:"default:false"`
}

type DeleteChatRequest struct {
	ID uint64 `json:"id"`
}

type ChatResponse struct {
	ID        uint64    `json:"id" bun:"primary_key"`
	Message   string    `json:"message"`
	Sender    string    `json:"sender"`
	CreatedAt time.Time `json:"created_at"`
}

type ChatList []Chat

func (c *Chat) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(
			&c.Message,
			validation.Required.Error("message is required"),
			validation.RuneLength(1, 255).Error("limited max 255 char"),
		),
	)
}

func (c *Chat) NewChatResponse() ChatResponse {
	return ChatResponse{
		ID:        c.ID,
		Message:   c.Message,
		Sender:    SenderMe,
		CreatedAt: c.CreatedAt,
	}
}

func (c *ChatList) NewChatResponse(userID string) []ChatResponse {
	res := []ChatResponse{}
	sender := ""
	for _, v := range *c {
		if v.UserID == userID {
			sender = SenderMe
		} else {
			sender = SenderOther
		}
		res = append(res, ChatResponse{
			ID:        v.ID,
			Message:   v.Message,
			Sender:    sender,
			CreatedAt: v.CreatedAt,
		})
	}
	return res
}
