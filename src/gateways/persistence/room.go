package persistence

import (
	"context"
	"homework/domain/model/room"
	"homework/domain/repository"
)

type Room struct{}

func NewRoom() *Room {
	return &Room{}
}

func (rr *Room) Create(db repository.DBConn) (string, error) {
	room := &room.Room{}
	_, err := db.NewInsert().Model(room).Exec(context.Background())
	if err != nil {
		return "", err
	}
	return room.ID, nil
}

func (rr *Room) CreateMap(db repository.DBConn, roomMap room.RoomMap) error {
	_, err := db.NewInsert().Model(&roomMap).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
