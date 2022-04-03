package repo

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(
	UserRepoProviderSet,
)

func AutoMigrate(db *gorm.DB) error {

	return db.AutoMigrate()
}
