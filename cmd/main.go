package main

import (
	"flag"
	"template/internal/app"
	"template/internal/app/config"
	"template/internal/app/router"
	"template/pkg/logx"
)

var (
	version = "1.0.0"
	name    = ""
	conf    string
)

func init() {
	flag.StringVar(&conf, "conf", "../configs/config.yaml", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()

	cfg, err := config.LoadConfig(conf)
	if err != nil {
		panic(err)
	}
	logger := logx.Zap(cfg)

	app, cleanup, err := initApp(cfg, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	if err = app.Run(cfg); err != nil {
		panic(err)
	}
}

func newApp(router *router.Router) *app.App {
	return app.New(
		app.SetName(name),
		app.SetVersion(version),
		app.SetEngine(router.RegisterAPI()),
	)
}
