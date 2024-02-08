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

func (rr *Room) Create(db repository.DBConn, r *room.Room) error {
	_, err := db.NewInsert().Model(r).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
