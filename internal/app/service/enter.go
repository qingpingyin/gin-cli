package service

import (
	"github.com/google/wire"
	"template/internal/app/service/user"
)

var ProviderSet = wire.NewSet(
	user.UserServiceProviderSet,
)
