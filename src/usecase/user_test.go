package usecase

import (
	"homework/config"
	userModel "homework/domain/model/user"
	"homework/domain/repository"
	"reflect"
	"testing"

	"github.com/uptrace/bun"
)

func Test_userUsecase_SignUp(t *testing.T) {
	type fields struct {
		ur repository.IUserRepository
		mu repository.Mail
		db *bun.DB
	}
	type args struct {
		user userModel.User
		cnf  config.Config
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    userModel.UserResponse
		want1   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := &userUsecase{
				ur: tt.fields.ur,
				mu: tt.fields.mu,
				db: tt.fields.db,
			}
			got, got1, err := uu.SignUp(tt.args.user, tt.args.cnf)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.SignUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.SignUp() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("userUsecase.SignUp() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
