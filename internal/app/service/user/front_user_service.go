package user

import (
	"go.uber.org/zap"
	"template/internal/app/repo/user"
)

type FrontUserService struct {
	UserRepo *user.FrontUserRepo
	Logger   *zap.Logger
}

func (us *FrontUserService) Login() {

}
