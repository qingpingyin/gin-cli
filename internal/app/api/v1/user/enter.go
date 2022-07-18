package user

import (
	"github.com/google/wire"
)

var UserAPIProviderSet = wire.NewSet(
	wire.Struct(new(FrontUserApi), "*"),
	wire.Struct(new(MerchantUserApi), "*"),
	wire.Struct(new(UserApi), "*"),
)

type UserApi struct {
	*FrontUserApi
	*MerchantUserApi
}
