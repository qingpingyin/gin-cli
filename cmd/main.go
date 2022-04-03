package main

import (
	"flag"
	"github.com/gin-gonic/gin"

	"gin-cli/internal/app"
	"gin-cli/internal/app/config"
)

var (
	Version = "1.0.0"
	Name    = ""
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
	app, cleanup, err := initApp(cfg)
	if err != nil {
		panic(err)
	}

	defer cleanup()
	if err = app.Run(cfg); err != nil {
		panic(err)
	}
}

func newApp(engine *gin.Engine) *app.App {
	return app.New(
		app.SetName(Name),
		app.SetVersion(Version),
		app.SetEngine(engine),
	)
}
