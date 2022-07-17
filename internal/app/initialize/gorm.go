package initialize

import (
	"gorm.io/gorm"
	"template/internal/app/config"
)

func InitGorm(c *config.Config) (*gorm.DB, func(), error) {

	//db, err := gormx.New(c)
	//if err != nil {
	//	return nil, nil, err
	//}
	//
	cleanFunc := func() {}
	//
	//if c.System.EnableAutoMigrate {
	//if err = model.AutoMigrate(db); err != nil {
	//	return nil, cleanFunc, err
	//}
	//是否启用自动迁移
	//}

	return nil, cleanFunc, nil
}
