package repo

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

var UserRepoProviderSet = wire.NewSet(wire.Struct(new(UserRepo), "*"))

type UserRepo struct {
	DB *gorm.DB
}

func (ur *UserRepo) GetUserInfo() {

}
