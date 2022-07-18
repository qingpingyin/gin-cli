package v1

import (
	"github.com/google/wire"
	"template/internal/app/api/v1/user"
)

var ProviderSet = wire.NewSet(
	user.UserAPIProviderSet,
)
