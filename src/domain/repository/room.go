package repository

import (
	"homework/domain/model/room"
)

//go:generate moq -pkg fakerepository -out ../../moq/fakerepository/room_test_moq.go . IRoomRepository
type IRoomRepository interface {
	Create(db DBConn, user *room.Room) error
}
