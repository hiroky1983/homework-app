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
		db *bun.DB
	}
	type args struct {
		user userModel.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    userModel.UserResponse
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := &userUsecase{
				ur: tt.fields.ur,
				db: tt.fields.db,
			}
			got, err := uu.SignUp(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.SignUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.SignUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_Login(t *testing.T) {
	type fields struct {
		ur repository.IUserRepository
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
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := &userUsecase{
				ur: tt.fields.ur,
				db: tt.fields.db,
			}
			got, err := uu.Login(tt.args.user, tt.args.cnf)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("userUsecase.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_LoginWithGoogle(t *testing.T) {
	type fields struct {
		ur repository.IUserRepository
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
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := &userUsecase{
				ur: tt.fields.ur,
				db: tt.fields.db,
			}
			got, err := uu.LoginWithGoogle(tt.args.user, tt.args.cnf)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.LoginWithGoogle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("userUsecase.LoginWithGoogle() = %v, want %v", got, tt.want)
			}
		})
	}
}
