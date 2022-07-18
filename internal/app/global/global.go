package global

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"template/internal/app/config"
)

var (
	HXSD_CONFIG config.Config
	HXSD_DB     *gorm.DB
	HXSD_LOG    *zap.Logger
	HXSD_REDIS  *redis.Client
)
