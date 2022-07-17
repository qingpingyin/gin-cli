package logx

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"template/internal/app/global"
	"template/internal/app/utils"
	"template/pkg/logx/internal"
)

func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.HXSD_CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.HXSD_CONFIG.Zap.Director)
		_ = os.Mkdir(global.HXSD_CONFIG.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.HXSD_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
