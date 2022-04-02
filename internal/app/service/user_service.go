package service

import (
	"github.com/google/wire"

	"gin-cli/internal/app/repo"
)

var UserServiceProviderSet = wire.NewSet(wire.Struct(new(UserService),"*"))



type UserService struct {
	UserRepo *repo.UserRepo
}


func (us *UserService)Login(){
	us.UserRepo.GetUserInfo()
}
