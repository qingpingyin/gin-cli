package repo

import (
	"context"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var UserRepoProviderSet = wire.NewSet(wire.Struct(new(UserRepo), "*"))

type UserRepo struct {
	DB *gorm.DB
}

func (ur *UserRepo) GetUserInfo() error {

	return nil

}
func (ur *UserRepo) UserIsExist(ctx context.Context, username, password string) (bool, error) {

	return true, nil

}
