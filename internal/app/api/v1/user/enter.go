package user

import "template/internal/app/service"

var (
	frontUserService    = service.FrontServiceGroupApp.FrontUserService
	merchantUserService = service.MerchantServiceGroupApp.MerchantUserService
)

type (
	FrontUserApiGroup struct {
		FrontUserApi
	}
	MerchantUserApiGroup struct {
		MerchantUserApi
	}
)
