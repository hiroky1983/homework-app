package repository

import (
	"context"
	"homework/domain/model/chat"
)

type IChatRepository interface {
	Create(db DBConn, user *chat.Chat) error
	ListChatByUserID(db DBConn, chatList *chat.ChatList) error
}

type chatRepository struct{}

func NewChatRepository() IChatRepository {
	return &chatRepository{}
}

func (cr *chatRepository) Create(db DBConn, chat *chat.Chat) error {
	_, err := db.NewInsert().Model(chat).Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (cr *chatRepository) ListChatByUserID(db DBConn, chatList *chat.ChatList) error {
	if err := db.NewSelect().Model((chatList)).Where("is_deleted = false").Scan(context.Background(), chatList); err != nil {
		return err
	}
	return nil
}
