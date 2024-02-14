// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package fakerepository

import (
	"homework/domain/model/chat"
	"homework/domain/repository"
	"sync"
)

// Ensure, that IChatRepositoryMock does implement repository.IChatRepository.
// If this is not the case, regenerate this file with moq.
var _ repository.IChatRepository = &IChatRepositoryMock{}

// IChatRepositoryMock is a mock implementation of repository.IChatRepository.
//
//	func TestSomethingThatUsesIChatRepository(t *testing.T) {
//
//		// make and configure a mocked repository.IChatRepository
//		mockedIChatRepository := &IChatRepositoryMock{
//			CreateFunc: func(db repository.DBConn, user *chat.Chat) error {
//				panic("mock out the Create method")
//			},
//			DeleteFunc: func(db repository.DBConn, chatID uint64) error {
//				panic("mock out the Delete method")
//			},
//			ListChatByUserIDFunc: func(db repository.DBConn, chatList *chat.ChatList, roomID string) error {
//				panic("mock out the ListChatByUserID method")
//			},
//		}
//
//		// use mockedIChatRepository in code that requires repository.IChatRepository
//		// and then make assertions.
//
//	}
type IChatRepositoryMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(db repository.DBConn, user *chat.Chat) error

	// DeleteFunc mocks the Delete method.
	DeleteFunc func(db repository.DBConn, chatID uint64) error

	// ListChatByUserIDFunc mocks the ListChatByUserID method.
	ListChatByUserIDFunc func(db repository.DBConn, chatList *chat.ChatList, roomID string) error

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Db is the db argument value.
			Db repository.DBConn
			// User is the user argument value.
			User *chat.Chat
		}
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// Db is the db argument value.
			Db repository.DBConn
			// ChatID is the chatID argument value.
			ChatID uint64
		}
		// ListChatByUserID holds details about calls to the ListChatByUserID method.
		ListChatByUserID []struct {
			// Db is the db argument value.
			Db repository.DBConn
			// ChatList is the chatList argument value.
			ChatList *chat.ChatList
			// RoomID is the roomID argument value.
			RoomID string
		}
	}
	lockCreate           sync.RWMutex
	lockDelete           sync.RWMutex
	lockListChatByUserID sync.RWMutex
}

// Create calls CreateFunc.
func (mock *IChatRepositoryMock) Create(db repository.DBConn, user *chat.Chat) error {
	if mock.CreateFunc == nil {
		panic("IChatRepositoryMock.CreateFunc: method is nil but IChatRepository.Create was just called")
	}
	callInfo := struct {
		Db   repository.DBConn
		User *chat.Chat
	}{
		Db:   db,
		User: user,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(db, user)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//
//	len(mockedIChatRepository.CreateCalls())
func (mock *IChatRepositoryMock) CreateCalls() []struct {
	Db   repository.DBConn
	User *chat.Chat
} {
	var calls []struct {
		Db   repository.DBConn
		User *chat.Chat
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// Delete calls DeleteFunc.
func (mock *IChatRepositoryMock) Delete(db repository.DBConn, chatID uint64) error {
	if mock.DeleteFunc == nil {
		panic("IChatRepositoryMock.DeleteFunc: method is nil but IChatRepository.Delete was just called")
	}
	callInfo := struct {
		Db     repository.DBConn
		ChatID uint64
	}{
		Db:     db,
		ChatID: chatID,
	}
	mock.lockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	mock.lockDelete.Unlock()
	return mock.DeleteFunc(db, chatID)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//
//	len(mockedIChatRepository.DeleteCalls())
func (mock *IChatRepositoryMock) DeleteCalls() []struct {
	Db     repository.DBConn
	ChatID uint64
} {
	var calls []struct {
		Db     repository.DBConn
		ChatID uint64
	}
	mock.lockDelete.RLock()
	calls = mock.calls.Delete
	mock.lockDelete.RUnlock()
	return calls
}

// ListChatByUserID calls ListChatByUserIDFunc.
func (mock *IChatRepositoryMock) ListChatByUserID(db repository.DBConn, chatList *chat.ChatList, roomID string) error {
	if mock.ListChatByUserIDFunc == nil {
		panic("IChatRepositoryMock.ListChatByUserIDFunc: method is nil but IChatRepository.ListChatByUserID was just called")
	}
	callInfo := struct {
		Db       repository.DBConn
		ChatList *chat.ChatList
		RoomID   string
	}{
		Db:       db,
		ChatList: chatList,
		RoomID:   roomID,
	}
	mock.lockListChatByUserID.Lock()
	mock.calls.ListChatByUserID = append(mock.calls.ListChatByUserID, callInfo)
	mock.lockListChatByUserID.Unlock()
	return mock.ListChatByUserIDFunc(db, chatList, roomID)
}

// ListChatByUserIDCalls gets all the calls that were made to ListChatByUserID.
// Check the length with:
//
//	len(mockedIChatRepository.ListChatByUserIDCalls())
func (mock *IChatRepositoryMock) ListChatByUserIDCalls() []struct {
	Db       repository.DBConn
	ChatList *chat.ChatList
	RoomID   string
} {
	var calls []struct {
		Db       repository.DBConn
		ChatList *chat.ChatList
		RoomID   string
	}
	mock.lockListChatByUserID.RLock()
	calls = mock.calls.ListChatByUserID
	mock.lockListChatByUserID.RUnlock()
	return calls
}
