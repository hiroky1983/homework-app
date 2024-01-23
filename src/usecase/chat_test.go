package usecase

import (
	"homework/domain/repository"
	"reflect"
	"testing"

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

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewChatUsecase(tt.args.ur, tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChatUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}
