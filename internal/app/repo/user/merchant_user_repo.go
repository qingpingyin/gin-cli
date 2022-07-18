package user

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type MerchantUserRepo struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func (us *MerchantUserRepo) Login() {}
