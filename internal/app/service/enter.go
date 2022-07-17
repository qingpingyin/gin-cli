package service

import "template/internal/app/service/user"

var (
	FrontServiceGroupApp    = new(FrontServiceGroup)
	MerchantServiceGroupApp = new(MerchantServiceGroup)
)

type (
	FrontServiceGroup struct {
		user.FrontUserService
	}

	MerchantServiceGroup struct {
		user.MerchantUserService
	}
)
