package usecase

import (
	"errors"
	chatModel "homework/domain/model/chat"
	"homework/domain/repository"
	"homework/moq/fakerepository"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func TestNewChatUsecase(t *testing.T) {
	type args struct {
		ur repository.IChatRepository
		db *bun.DB
	}
	tests := []struct {
		name string
		args args
		want IChatUsecase
	}{
		{
			name: "success",
			args: args{
				ur: nil,
				db: nil,
			},
			want: &chatUsecase{
				ur: nil,
				db: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewChatUsecase(tt.args.ur, tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChatUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_chatUsecase_Create(t *testing.T) {
	now := time.Now()
	type fields struct {
		ur repository.IChatRepository
		db *bun.DB
	}
	type args struct {
		c chatModel.Chat
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    chatModel.ChatResponse
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				ur: &fakerepository.IChatRepositoryMock{
					CreateFunc: func(db repository.DBConn, user *chatModel.Chat) error {
						return nil
					},
				},
			},
			args: args{
				c: chatModel.Chat{
					ID:        1,
					UserID:    uuid.New().String(),
					Message:   "test",
					CreatedAt: now,
				},
			},
			want: chatModel.ChatResponse{
				ID:        1,
				Sender:    chatModel.SenderMe,
				Message:   "test",
				CreatedAt: now,
			},
			wantErr: false,
		},
		{
			name: "failer",
			fields: fields{
				ur: &fakerepository.IChatRepositoryMock{
					CreateFunc: func(db repository.DBConn, user *chatModel.Chat) error {
						return errors.New("error")
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cu := &chatUsecase{
				ur: tt.fields.ur,
				db: tt.fields.db,
			}
			got, err := cu.Create(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("chatUsecase.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("chatUsecase.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_chatUsecase_Delete(t *testing.T) {
	type fields struct {
		ur repository.IChatRepository
		db *bun.DB
	}
	type args struct {
		c uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				ur: &fakerepository.IChatRepositoryMock{
					DeleteFunc: func(db repository.DBConn, chatID uint64) error {
						return nil
					},
				},
			},
			args: args{
				c: 1,
			},
			wantErr: false,
		},
		{
			name: "failer",
			fields: fields{
				ur: &fakerepository.IChatRepositoryMock{
					DeleteFunc: func(db repository.DBConn, chatID uint64) error {
						return errors.New("error")
					},
				},
			},
			args: args{
				c: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cu := &chatUsecase{
				ur: tt.fields.ur,
				db: tt.fields.db,
			}
			if err := cu.Delete(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("chatUsecase.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_chatUsecase_List(t *testing.T) {
	id := uuid.New().String()
	now := time.Now()
	type fields struct {
		ur repository.IChatRepository
		db *bun.DB
	}
	type args struct {
		userID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []chatModel.ChatResponse
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				ur: &fakerepository.IChatRepositoryMock{
					ListChatByUserIDFunc: func(db repository.DBConn, chatList *chatModel.ChatList) error {
						*chatList = append(*chatList, chatModel.Chat{
							ID:        1,
							UserID:    id,
							Message:   "test",
							CreatedAt: now,
						})
						return nil
					},
				},
			},
			args: args{
				userID: id,
			},
			want: []chatModel.ChatResponse{
				{
					ID:        1,
					Sender:    chatModel.SenderMe,
					Message:   "test",
					CreatedAt: now,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cu := &chatUsecase{
				ur: tt.fields.ur,
				db: tt.fields.db,
			}
			got, err := cu.List(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("chatUsecase.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("chatUsecase.List() = %v, want %v", got, tt.want)
			}
		})
	}
}
