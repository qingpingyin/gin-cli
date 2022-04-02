package repo

import (
	"github.com/google/wire"
)

var UserRepoProviderSet = wire.NewSet(wire.Struct(new(UserRepo),"*"))

type UserRepo struct {
	// db *gorm.DB
}


func (ur *UserRepo)GetUserInfo(){

}