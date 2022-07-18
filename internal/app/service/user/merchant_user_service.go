package user

import (
	"go.uber.org/zap"
	"template/internal/app/repo/user"
)

type MerchantUserService struct {
	UserRepo *user.MerchantUserRepo
	Logger   *zap.Logger
}

func (us *MerchantUserService) Login() {

}
