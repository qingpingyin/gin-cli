package global

import (
	"go.uber.org/zap"
	"template/internal/app/config"
)

var (
	HXSD_CONFIG config.Config
	//HXSD_DB
	HXSD_LOG *zap.Logger
)
