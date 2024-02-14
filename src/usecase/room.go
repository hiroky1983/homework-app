package usecase

import (
	"homework/domain/model/room"
	"homework/domain/repository"

	"github.com/uptrace/bun"
)

type IRoomUsecase interface {
	Create(userID string, otherUser room.RoomMap) (string, error)
}

type roomUsecase struct {
	rr repository.IRoomRepository
	db *bun.DB
}

func NewRoomUsecase(rr repository.IRoomRepository, db *bun.DB) IRoomUsecase {
	return &roomUsecase{rr, db}
}

func (ru *roomUsecase) Create(userID string, otherUser room.RoomMap) (string, error) {
	tx, err := ru.db.Begin()
	if err != nil {
		return "", err
	}
	RoomID, err := ru.rr.Create(tx)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	roomMapMe := room.RoomMap{
		UserID: userID,
		RoomID: RoomID,
	}

	roomMapOther := room.RoomMap{
		UserID: otherUser.UserID,
		RoomID: RoomID,
	}
	if err := ru.rr.CreateMap(tx, roomMapMe); err != nil {
		tx.Rollback()
		return "", err
	}

	if err := ru.rr.CreateMap(tx, roomMapOther); err != nil {
		tx.Rollback()
		return "", err
	}
	tx.Commit()
	return RoomID, nil
}
