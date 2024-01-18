package usecase

import (
	"homework/domain/repository"
)

type IChatUsecase interface {}

type chatUsecase struct {
	ur repository.IUserRepository
}

func NewChatUsecase(ur repository.IUserRepository) IChatUsecase {
	return &chatUsecase{ur}
}