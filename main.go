package main

import (
	"flag"
	"go.uber.org/zap"
	"template/internal/app"
	"template/internal/app/config"
	"template/internal/app/global"
	"template/internal/app/initialize"
	"template/pkg/logx"
)

var (
	conf    string
	version = "hxsh-1.0"
	name    = "hxsd"
)

func init() {
	flag.StringVar(&conf, "conf", "./configs/config.yaml", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()

	cfg, err := config.LoadConfig(conf)
	if err != nil {
		panic(err)
	}
	global.HXSD_CONFIG = *cfg
	global.HXSD_LOG = logx.Zap()
	zap.ReplaceGlobals(global.HXSD_LOG)
	db, cancel, err := initialize.InitGorm(cfg)
	if err != nil {
		panic(err)
	}
	cancel()
	global.HXSD_DB = db
	initialize.Redis(cfg)
	err = app.New(
		app.SetVersion(version),
		app.SetName(name),
		app.SetEngine(initialize.Routers()),
	).Run(cfg)
	if err != nil {
		panic(err)
	}
}
