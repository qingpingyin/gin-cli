package user

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type FrontUserRepo struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func (us *FrontUserRepo) Login() {
}
