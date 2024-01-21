package usecase

import (
	chatModel "homework/domain/model/chat"
	"homework/domain/repository"

	"github.com/uptrace/bun"
)

type IChatUsecase interface {
	Create(chat chatModel.Chat) (chatModel.ChatResponse, error)
}

type chatUsecase struct {
	ur repository.IChatRepository
	db *bun.DB
}

func NewChatUsecase(ur repository.IChatRepository, db *bun.DB) IChatUsecase {
	return &chatUsecase{ur, db}
}

func (cu *chatUsecase) Create(c chatModel.Chat) (chatModel.ChatResponse, error) {
	if err := c.Validate(); err != nil {
		return chatModel.ChatResponse{}, err
	}

	if err := cu.ur.Create(cu.db, &c); err != nil {
		return chatModel.ChatResponse{}, err
	}
	res := c.NewChatResponse()
	
	return res, nil
}
