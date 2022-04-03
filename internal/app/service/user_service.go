package service

import (
	"context"
	"gin-cli/internal/app/model/request"
	"gin-cli/third_party/serror"
	"github.com/google/wire"

	"gin-cli/internal/app/repo"
)

var UserServiceProviderSet = wire.NewSet(wire.Struct(new(UserService), "*"))

type UserService struct {
	UserRepo *repo.UserRepo
}

func (us *UserService) Login(ctx context.Context, loginReq request.LoginRequest) *serror.ResponseError {
	exist, err := us.UserRepo.UserIsExist(ctx, loginReq.UserName, loginReq.Password)
	if err != nil {
		//记录日志
		//返回业务错误
		return serror.ErrInternalServer
	}
	if !exist {
		return serror.ErrUserNotExist
	}

	return nil
}
