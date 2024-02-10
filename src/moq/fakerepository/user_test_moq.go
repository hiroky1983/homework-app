// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package fakerepository

import (
	"homework/domain/model/user"
	"homework/domain/repository"
	"sync"
)

// Ensure, that IUserRepositoryMock does implement repository.IUserRepository.
// If this is not the case, regenerate this file with moq.
var _ repository.IUserRepository = &IUserRepositoryMock{}

// IUserRepositoryMock is a mock implementation of repository.IUserRepository.
//
//	func TestSomethingThatUsesIUserRepository(t *testing.T) {
//
//		// make and configure a mocked repository.IUserRepository
//		mockedIUserRepository := &IUserRepositoryMock{
//			CreateUserFunc: func(db repository.DBConn, userMoqParam *user.User) error {
//				panic("mock out the CreateUser method")
//			},
//			GetProfileFunc: func(db repository.DBConn, u *user.User, userID string) error {
//				panic("mock out the GetProfile method")
//			},
//			GetUserByEmailFunc: func(db repository.DBConn, userMoqParam *user.User, email string) error {
//				panic("mock out the GetUserByEmail method")
//			},
//			GetUserByIDFunc: func(db repository.DBConn, userMoqParam *user.User, userID string) error {
//				panic("mock out the GetUserByID method")
//			},
//			ListUserFunc: func(db repository.DBConn, u *[]user.User, userID string) error {
//				panic("mock out the ListUser method")
//			},
//			UpdateIsVerifiedUserFunc: func(db repository.DBConn, userID string) error {
//				panic("mock out the UpdateIsVerifiedUser method")
//			},
//			UpdateUserFunc: func(db repository.DBConn, u *user.User) error {
//				panic("mock out the UpdateUser method")
//			},
//		}
//
//		// use mockedIUserRepository in code that requires repository.IUserRepository
//		// and then make assertions.
//
//	}
type IUserRepositoryMock struct {
	// CreateUserFunc mocks the CreateUser method.
	CreateUserFunc func(db repository.DBConn, userMoqParam *user.User) error

	// GetProfileFunc mocks the GetProfile method.
	GetProfileFunc func(db repository.DBConn, u *user.User, userID string) error

	// GetUserByEmailFunc mocks the GetUserByEmail method.
	GetUserByEmailFunc func(db repository.DBConn, userMoqParam *user.User, email string) error

	// GetUserByIDFunc mocks the GetUserByID method.
	GetUserByIDFunc func(db repository.DBConn, userMoqParam *user.User, userID string) error

	// ListUserFunc mocks the ListUser method.
	ListUserFunc func(db repository.DBConn, u *[]user.User, userID string) error

	// UpdateIsVerifiedUserFunc mocks the UpdateIsVerifiedUser method.
	UpdateIsVerifiedUserFunc func(db repository.DBConn, userID string) error

	// UpdateUserFunc mocks the UpdateUser method.
	UpdateUserFunc func(db repository.DBConn, u *user.User) error

	// calls tracks calls to the methods.
	calls struct {
		// CreateUser holds details about calls to the CreateUser method.
		CreateUser []struct {
			// Db is the db argument value.
			Db repository.DBConn
			// UserMoqParam is the userMoqParam argument value.
			UserMoqParam *user.User
		}
		// GetProfile holds details about calls to the GetProfile method.
		GetProfile []struct {
			// Db is the db argument value.
			Db repository.DBConn
			// U is the u argument value.
			U *user.User
			// UserID is the userID argument value.
			UserID string
		}
		// GetUserByEmail holds details about calls to the GetUserByEmail method.
		GetUserByEmail []struct {
			// Db is the db argument value.
			Db repository.DBConn
			// UserMoqParam is the userMoqParam argument value.
			UserMoqParam *user.User
			// Email is the email argument value.
			Email string
		}
		// GetUserByID holds details about calls to the GetUserByID method.
		GetUserByID []struct {
			// Db is the db argument value.
			Db repository.DBConn
			// UserMoqParam is the userMoqParam argument value.
			UserMoqParam *user.User
			// UserID is the userID argument value.
			UserID string
		}
		// ListUser holds details about calls to the ListUser method.
		ListUser []struct {
			// Db is the db argument value.
			Db repository.DBConn
			// U is the u argument value.
			U *[]user.User
			// UserID is the userID argument value.
			UserID string
		}
		// UpdateIsVerifiedUser holds details about calls to the UpdateIsVerifiedUser method.
		UpdateIsVerifiedUser []struct {
			// Db is the db argument value.
			Db repository.DBConn
			// UserID is the userID argument value.
			UserID string
		}
		// UpdateUser holds details about calls to the UpdateUser method.
		UpdateUser []struct {
			// Db is the db argument value.
			Db repository.DBConn
			// U is the u argument value.
			U *user.User
		}
	}
	lockCreateUser           sync.RWMutex
	lockGetProfile           sync.RWMutex
	lockGetUserByEmail       sync.RWMutex
	lockGetUserByID          sync.RWMutex
	lockListUser             sync.RWMutex
	lockUpdateIsVerifiedUser sync.RWMutex
	lockUpdateUser           sync.RWMutex
}

// CreateUser calls CreateUserFunc.
func (mock *IUserRepositoryMock) CreateUser(db repository.DBConn, userMoqParam *user.User) error {
	if mock.CreateUserFunc == nil {
		panic("IUserRepositoryMock.CreateUserFunc: method is nil but IUserRepository.CreateUser was just called")
	}
	callInfo := struct {
		Db           repository.DBConn
		UserMoqParam *user.User
	}{
		Db:           db,
		UserMoqParam: userMoqParam,
	}
	mock.lockCreateUser.Lock()
	mock.calls.CreateUser = append(mock.calls.CreateUser, callInfo)
	mock.lockCreateUser.Unlock()
	return mock.CreateUserFunc(db, userMoqParam)
}

// CreateUserCalls gets all the calls that were made to CreateUser.
// Check the length with:
//
//	len(mockedIUserRepository.CreateUserCalls())
func (mock *IUserRepositoryMock) CreateUserCalls() []struct {
	Db           repository.DBConn
	UserMoqParam *user.User
} {
	var calls []struct {
		Db           repository.DBConn
		UserMoqParam *user.User
	}
	mock.lockCreateUser.RLock()
	calls = mock.calls.CreateUser
	mock.lockCreateUser.RUnlock()
	return calls
}

// GetProfile calls GetProfileFunc.
func (mock *IUserRepositoryMock) GetProfile(db repository.DBConn, u *user.User, userID string) error {
	if mock.GetProfileFunc == nil {
		panic("IUserRepositoryMock.GetProfileFunc: method is nil but IUserRepository.GetProfile was just called")
	}
	callInfo := struct {
		Db     repository.DBConn
		U      *user.User
		UserID string
	}{
		Db:     db,
		U:      u,
		UserID: userID,
	}
	mock.lockGetProfile.Lock()
	mock.calls.GetProfile = append(mock.calls.GetProfile, callInfo)
	mock.lockGetProfile.Unlock()
	return mock.GetProfileFunc(db, u, userID)
}

// GetProfileCalls gets all the calls that were made to GetProfile.
// Check the length with:
//
//	len(mockedIUserRepository.GetProfileCalls())
func (mock *IUserRepositoryMock) GetProfileCalls() []struct {
	Db     repository.DBConn
	U      *user.User
	UserID string
} {
	var calls []struct {
		Db     repository.DBConn
		U      *user.User
		UserID string
	}
	mock.lockGetProfile.RLock()
	calls = mock.calls.GetProfile
	mock.lockGetProfile.RUnlock()
	return calls
}

// GetUserByEmail calls GetUserByEmailFunc.
func (mock *IUserRepositoryMock) GetUserByEmail(db repository.DBConn, userMoqParam *user.User, email string) error {
	if mock.GetUserByEmailFunc == nil {
		panic("IUserRepositoryMock.GetUserByEmailFunc: method is nil but IUserRepository.GetUserByEmail was just called")
	}
	callInfo := struct {
		Db           repository.DBConn
		UserMoqParam *user.User
		Email        string
	}{
		Db:           db,
		UserMoqParam: userMoqParam,
		Email:        email,
	}
	mock.lockGetUserByEmail.Lock()
	mock.calls.GetUserByEmail = append(mock.calls.GetUserByEmail, callInfo)
	mock.lockGetUserByEmail.Unlock()
	return mock.GetUserByEmailFunc(db, userMoqParam, email)
}

// GetUserByEmailCalls gets all the calls that were made to GetUserByEmail.
// Check the length with:
//
//	len(mockedIUserRepository.GetUserByEmailCalls())
func (mock *IUserRepositoryMock) GetUserByEmailCalls() []struct {
	Db           repository.DBConn
	UserMoqParam *user.User
	Email        string
} {
	var calls []struct {
		Db           repository.DBConn
		UserMoqParam *user.User
		Email        string
	}
	mock.lockGetUserByEmail.RLock()
	calls = mock.calls.GetUserByEmail
	mock.lockGetUserByEmail.RUnlock()
	return calls
}

// GetUserByID calls GetUserByIDFunc.
func (mock *IUserRepositoryMock) GetUserByID(db repository.DBConn, userMoqParam *user.User, userID string) error {
	if mock.GetUserByIDFunc == nil {
		panic("IUserRepositoryMock.GetUserByIDFunc: method is nil but IUserRepository.GetUserByID was just called")
	}
	callInfo := struct {
		Db           repository.DBConn
		UserMoqParam *user.User
		UserID       string
	}{
		Db:           db,
		UserMoqParam: userMoqParam,
		UserID:       userID,
	}
	mock.lockGetUserByID.Lock()
	mock.calls.GetUserByID = append(mock.calls.GetUserByID, callInfo)
	mock.lockGetUserByID.Unlock()
	return mock.GetUserByIDFunc(db, userMoqParam, userID)
}

// GetUserByIDCalls gets all the calls that were made to GetUserByID.
// Check the length with:
//
//	len(mockedIUserRepository.GetUserByIDCalls())
func (mock *IUserRepositoryMock) GetUserByIDCalls() []struct {
	Db           repository.DBConn
	UserMoqParam *user.User
	UserID       string
} {
	var calls []struct {
		Db           repository.DBConn
		UserMoqParam *user.User
		UserID       string
	}
	mock.lockGetUserByID.RLock()
	calls = mock.calls.GetUserByID
	mock.lockGetUserByID.RUnlock()
	return calls
}

// ListUser calls ListUserFunc.
func (mock *IUserRepositoryMock) ListUser(db repository.DBConn, u *[]user.User, userID string) error {
	if mock.ListUserFunc == nil {
		panic("IUserRepositoryMock.ListUserFunc: method is nil but IUserRepository.ListUser was just called")
	}
	callInfo := struct {
		Db     repository.DBConn
		U      *[]user.User
		UserID string
	}{
		Db:     db,
		U:      u,
		UserID: userID,
	}
	mock.lockListUser.Lock()
	mock.calls.ListUser = append(mock.calls.ListUser, callInfo)
	mock.lockListUser.Unlock()
	return mock.ListUserFunc(db, u, userID)
}

// ListUserCalls gets all the calls that were made to ListUser.
// Check the length with:
//
//	len(mockedIUserRepository.ListUserCalls())
func (mock *IUserRepositoryMock) ListUserCalls() []struct {
	Db     repository.DBConn
	U      *[]user.User
	UserID string
} {
	var calls []struct {
		Db     repository.DBConn
		U      *[]user.User
		UserID string
	}
	mock.lockListUser.RLock()
	calls = mock.calls.ListUser
	mock.lockListUser.RUnlock()
	return calls
}

// UpdateIsVerifiedUser calls UpdateIsVerifiedUserFunc.
func (mock *IUserRepositoryMock) UpdateIsVerifiedUser(db repository.DBConn, userID string) error {
	if mock.UpdateIsVerifiedUserFunc == nil {
		panic("IUserRepositoryMock.UpdateIsVerifiedUserFunc: method is nil but IUserRepository.UpdateIsVerifiedUser was just called")
	}
	callInfo := struct {
		Db     repository.DBConn
		UserID string
	}{
		Db:     db,
		UserID: userID,
	}
	mock.lockUpdateIsVerifiedUser.Lock()
	mock.calls.UpdateIsVerifiedUser = append(mock.calls.UpdateIsVerifiedUser, callInfo)
	mock.lockUpdateIsVerifiedUser.Unlock()
	return mock.UpdateIsVerifiedUserFunc(db, userID)
}

// UpdateIsVerifiedUserCalls gets all the calls that were made to UpdateIsVerifiedUser.
// Check the length with:
//
//	len(mockedIUserRepository.UpdateIsVerifiedUserCalls())
func (mock *IUserRepositoryMock) UpdateIsVerifiedUserCalls() []struct {
	Db     repository.DBConn
	UserID string
} {
	var calls []struct {
		Db     repository.DBConn
		UserID string
	}
	mock.lockUpdateIsVerifiedUser.RLock()
	calls = mock.calls.UpdateIsVerifiedUser
	mock.lockUpdateIsVerifiedUser.RUnlock()
	return calls
}

// UpdateUser calls UpdateUserFunc.
func (mock *IUserRepositoryMock) UpdateUser(db repository.DBConn, u *user.User) error {
	if mock.UpdateUserFunc == nil {
		panic("IUserRepositoryMock.UpdateUserFunc: method is nil but IUserRepository.UpdateUser was just called")
	}
	callInfo := struct {
		Db repository.DBConn
		U  *user.User
	}{
		Db: db,
		U:  u,
	}
	mock.lockUpdateUser.Lock()
	mock.calls.UpdateUser = append(mock.calls.UpdateUser, callInfo)
	mock.lockUpdateUser.Unlock()
	return mock.UpdateUserFunc(db, u)
}

// UpdateUserCalls gets all the calls that were made to UpdateUser.
// Check the length with:
//
//	len(mockedIUserRepository.UpdateUserCalls())
func (mock *IUserRepositoryMock) UpdateUserCalls() []struct {
	Db repository.DBConn
	U  *user.User
} {
	var calls []struct {
		Db repository.DBConn
		U  *user.User
	}
	mock.lockUpdateUser.RLock()
	calls = mock.calls.UpdateUser
	mock.lockUpdateUser.RUnlock()
	return calls
}
