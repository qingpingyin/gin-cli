package main

import (
	"flag"

	"github.com/gin-gonic/gin"

	"gin-cli/internal/app"
)

var (
	Version = "1.0.0"
	Name    = ""
	conf    string
)

func init() {
	flag.StringVar(&conf, "conf", "../internal/configs", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()

	app, cleanup, err := initApp()
	if err != nil {
		panic(err)
	}

	defer cleanup()
	if err = app.Run();err !=nil {
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
