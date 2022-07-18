package initialize

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"template/internal/app/config"
	"template/internal/app/global"
)

func Redis(cfg *config.Config) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password, // no password set
		DB:       cfg.Redis.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.HXSD_LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.HXSD_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.HXSD_REDIS = client
	}
}
