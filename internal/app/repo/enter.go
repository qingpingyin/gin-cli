package repo

import (
	"github.com/google/wire"
	"template/internal/app/repo/user"
)

var ProviderSet = wire.NewSet(
	user.UserRepoProviderSet,
)
