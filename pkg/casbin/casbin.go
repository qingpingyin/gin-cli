package casbin

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"os"
)

var (
	Enforcer *casbin.Enforcer
)

// 初始化 casbin
func InitCasbin(db *gorm.DB) (*casbin.Enforcer, error) {

	// 判断是否有缓存
	if Enforcer != nil {
		return Enforcer, nil
	}
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		return nil, err
	}

	casbinPath := viper.GetString("casbin.path")
	if casbinPath == "" {
		wd, _ := os.Getwd()
		casbinPath = wd + "/conf/rbac_model.conf"
	}
	Enforcer, err = casbin.NewEnforcer(casbinPath, adapter)
	if err != nil {
		return nil, err
	}

	if err = Enforcer.LoadPolicy(); err != nil {
		return nil, err
	}

	return Enforcer, nil
}

// 清空缓存
func ClearEnforcer() {
	Enforcer = nil
}
