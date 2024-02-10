package usecase

import (
	"homework/domain/model/room"
	"homework/domain/repository"

	"github.com/uptrace/bun"
)

type IRoomUsecase interface {
	Create(userID string) error
}

type roomUsecase struct {
	rr repository.IRoomRepository
	db *bun.DB
}

func NewRoomUsecase(rr repository.IRoomRepository, db *bun.DB) IRoomUsecase {
	return &roomUsecase{rr, db}
}

func (ru *roomUsecase) Create(userID string) error {
	tx, err := ru.db.Begin()
	if err != nil {
		return err
	}
	ID, err := ru.rr.Create(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	roomMap := room.RoomMap{
		UserID: userID,
		RoomID: ID,
	}
	if err := ru.rr.CreateMap(tx, roomMap); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
