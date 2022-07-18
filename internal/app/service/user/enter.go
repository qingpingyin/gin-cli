package user

import "github.com/google/wire"

var UserServiceProviderSet = wire.NewSet(
	wire.Struct(new(FrontUserService), "*"),
	wire.Struct(new(MerchantUserService), "*"),
)
