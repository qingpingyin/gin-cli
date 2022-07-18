package gormx

import (
	"gorm.io/gorm"
	"template/internal/app/config"
)

var Gorm = new(_gorm)

type _gorm struct{}

// Config gorm 自定义配置
// Author [SliverHorn](https://github.com/SliverHorn)
func (g *_gorm) Config(cfg *config.Config) *gorm.Config {
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	//_default := logger.New(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
	//	SlowThreshold: 200 * time.Millisecond,
	//	LogLevel:      logger.Warn,
	//	Colorful:      true,
	//})

	//switch cfg.Mysql.LogMode {
	//case "silent", "Silent":
	//	config.Logger = _default.LogMode(logger.Silent)
	//case "error", "Error":
	//	config.Logger = _default.LogMode(logger.Error)
	//case "warn", "Warn":
	//	config.Logger = _default.LogMode(logger.Warn)
	//case "info", "Info":
	//	config.Logger = _default.LogMode(logger.Info)
	//default:
	//	config.Logger = _default.LogMode(logger.Info)
	//}
	return config
}
