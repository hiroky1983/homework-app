package repository

import "homework/domain/model/room"

//go:generate moq -pkg fakerepository -out ../../moq/fakerepository/room_test_moq.go . IRoomRepository
type IRoomRepository interface {
	Create(db DBConn) (string, error)
	CreateMap(db DBConn, roomMap room.RoomMap) error
	GetRoomByUserID(db DBConn, roomMap *room.RoomMap, userID string) error
	GetUserIDByRoomID(db DBConn, roomMap []*room.RoomMap, roomID string) ([]string, error)
}
