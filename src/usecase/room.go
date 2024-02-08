package usecase

import (
	"homework/config"
	"homework/domain/model/room"

	"github.com/uptrace/bun"
)

type IRoomUsecase interface {
	Create(r room.Room, conf config.Config) (room.Room, error)
}

type roomUsecase struct {
	db *bun.DB
}

func NewRoomUsecase(db *bun.DB) IRoomUsecase {
	return &roomUsecase{db}
}

func (ru *roomUsecase) Create(r room.Room, cnf config.Config) (room.Room, error) {
	tx, err := ru.db.Begin()
	if err != nil {
		return room.Room{}, err
	}
	// if err := ru.ur.CreateRoom(tx, &r); err != nil {
	// 	tx.Rollback()
	// 	return room.Room{}, err
	// }
	tx.Commit()
	return r, nil
}
