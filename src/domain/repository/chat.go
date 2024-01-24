package repository

import (
	"homework/domain/model/chat"
)

//go:generate moq -pkg fakerepository -out ../../moq/fakerepository/chat_test_moq.go . IChatRepository
type IChatRepository interface {
	Create(db DBConn, user *chat.Chat) error
	ListChatByUserID(db DBConn, chatList *chat.ChatList) error
	Delete(db DBConn, chatID uint64) error
}
