package v1

import "template/internal/app/api/v1/user"

var (
	FrontApiGroupApp    = new(FrontApiGroup)
	MerchantApiGroupApp = new(MerchantApiGroup)
)

type FrontApiGroup struct {
	UserApiGroup user.FrontUserApiGroup
}

type MerchantApiGroup struct {
	UserApiGroup user.MerchantUserApiGroup
}
