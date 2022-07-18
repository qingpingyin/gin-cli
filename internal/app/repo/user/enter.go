package user

import "github.com/google/wire"

var UserRepoProviderSet = wire.NewSet(
	wire.Struct(new(FrontUserRepo), "*"),
	wire.Struct(new(MerchantUserRepo), "*"),
)
