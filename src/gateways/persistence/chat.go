package persistence

import (
	"context"
	"homework/domain/model/chat"
	"homework/domain/repository"
)

type Chat struct{}

func NewChatRepository() *Chat {
	return &Chat{}
}

func (cr *Chat) Create(db repository.DBConn, chat *chat.Chat) error {
	_, err := db.NewInsert().Model(chat).Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (cr *Chat) ListChatByUserID(db repository.DBConn, chatList *chat.ChatList) error {
	if err := db.NewSelect().Model((chatList)).Where("is_deleted = false").Scan(context.Background(), chatList); err != nil {
		return err
	}
	return nil
}

func (cr *Chat) Delete(db repository.DBConn, chatID uint64) error {
	_, err := db.NewUpdate().Model(&chat.Chat{}).Set("is_deleted = true").Where("id = ?", chatID).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
