package logx

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"template/internal/app/config"
	"template/internal/app/utils"
)

func Zap(cfg *config.Config) (logger *zap.Logger) {
	if ok, _ := utils.PathExists(cfg.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", cfg.Zap.Director)
		_ = os.Mkdir(cfg.Director, os.ModePerm)
	}

	//cores := internal.Zap.GetZapCores()
	//logger = zap.New(zapcore.NewTee(cores...))
	//
	//if cfg.Zap.ShowLine {
	//	logger = logger.WithOptions(zap.AddCaller())
	//}
	return logger
}
